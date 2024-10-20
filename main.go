package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
	// "github.com/joho/godotenv"
	// "sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
)

type Config struct {
	BotToken     string   `json:"bot_token"`
	AllowedUsers []int64  `json:"allowed_users"`
	Scripts []Script
}

type Script struct{

	ID        string   `json:id`
	Command   string   `json:command`
 	Comment   string   `json:comment`

}

type Bot struct {
	api            *tgbotapi.BotAPI
	allowedUserIDs map[int64]bool
	logFile        string
	workingDir     string 
	sudoPassword   string
	config         Config
}

func NewBot(config Config, logFile string) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(config.BotToken)
	if err != nil {
		return nil, fmt.Errorf("failed to create bot: %v", err)
	}

	allowedMap := make(map[int64]bool)
	for _, userID := range config.AllowedUsers {
		allowedMap[userID] = true
	}

	return &Bot{
		api:            api,
		allowedUserIDs: allowedMap,
		logFile:        logFile,
		workingDir:     os.Getenv("."), 
		config:         config,           
	}, nil
}

func (b *Bot) isAuthorized(userID int64) bool {
	return b.allowedUserIDs[userID]
	// return true
}

func printBanner() {
	myFigure := figure.NewFigure("Termigram", "doom", true)
	
	color.Cyan(myFigure.String())
	
	color.Yellow("Version: 1.0.0")
	color.Yellow("Remote Command Execution Bot")
	fmt.Println()
	
	color.Green("• Telegram Bot Started")
	color.Green("• Logs will be saved to: bot_commands.log")
	color.Green("• Configuration loaded from: config.json")
	fmt.Println()
	
	// Print warning
	color.Red("⚠ Warning: This bot provides system-level access. Ensure proper security measures.")
	fmt.Println()
}

func (b *Bot) logCommand(userID int64, command string) error {
	f, err := os.OpenFile(b.logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	logEntry := fmt.Sprintf("[%s] UserID: %d, Command: %s\n",
		time.Now().Format(time.RFC3339),
		userID,
		command)

	if _, err := f.WriteString(logEntry); err != nil {
		return err
	}
	return nil
}

func (b *Bot) executeCommand(command string) (string, error) {
	var cmd *exec.Cmd
	var output []byte
	var err error

	
	if strings.HasPrefix(command, "cd ") {
		dir := strings.TrimPrefix(command, "cd ")
		if err := os.Chdir(dir); err != nil {
			return "", fmt.Errorf("failed to change directory: %v", err)
		}
		b.workingDir, _ = os.Getwd() 
		return fmt.Sprintf("Changed directory to: %s", b.workingDir), nil
	} else {
		cmd = exec.Command("/bin/sh", "-c", command)
		cmd.Dir = b.workingDir 
	}
	

	output, err = cmd.CombinedOutput()
	if err != nil {
		return string(output), fmt.Errorf("command execution failed: %v", err)
	}
	return string(output), nil
}

func (b *Bot) handleStart(message *tgbotapi.Message) {
	response := "Welcome to the Remote Command Execution Bot!\n" +
		"Available commands:\n" +
		"/start - Show this help message\n" +
		"/getlogs - Get command execution logs\n" +
		"/listscripts - Get the list of scripts defined\n"+
		"/runscript script_name - to run a defined script \n"+
		"For regular commands, just type the command\n" +
		fmt.Sprintf("Your Telegram ID: %d", message.From.ID)

	msg := tgbotapi.NewMessage(message.Chat.ID, response)
	b.api.Send(msg)
}

func (b *Bot) handleGetLogs(message *tgbotapi.Message) {

	if !b.isAuthorized(message.From.ID) {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Unauthorized access")
		b.api.Send(msg)
		return
	}
	logs, err := os.ReadFile(b.logFile)
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Failed to read logs")
		b.api.Send(msg)
		return
	}

	response := "Recent command logs:\n" + string(logs)
	msg := tgbotapi.NewMessage(message.Chat.ID, response)
	b.api.Send(msg)
}


func (b *Bot) listScripts(message *tgbotapi.Message){
	if !b.isAuthorized(message.From.ID) {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Unauthorized access")
		b.api.Send(msg)
		return
	}

	response := "";
	for _, script := range b.config.Scripts{
		response += script.ID +"   "+ script.Comment + "    " + script.Command
		response += "\n"

	}
	msg := tgbotapi.NewMessage(message.Chat.ID, response)
	b.api.Send(msg)

}

func (b *Bot) handleScript(message *tgbotapi.Message) {
	if !b.isAuthorized(message.From.ID) {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Unauthorized access")
		b.api.Send(msg)
		return
	}

	scriptID := strings.TrimPrefix(message.Text, "/runscript ") 
	var scriptToRun Script
	found := false
	for _, s := range b.config.Scripts {
		if s.ID == scriptID {
			scriptToRun = s
			found = true
			break
		}
	}

	if !found {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Script not found")
		b.api.Send(msg)
		return
	}

	// Execute the script
	Command := "./" + scriptToRun.Command
	output, err := b.executeCommand(Command)
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Error running script: %v\n%s", err, output))
		b.api.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Script output:\n%s", output))
	b.api.Send(msg)
}

func (b *Bot) handleCommand(message *tgbotapi.Message) {
	command := message.Text
	isSudo := strings.HasPrefix(command, "sudo ")
	// Log the command
	if err := b.logCommand(message.From.ID, command); err != nil {
		log.Printf("Failed to log command: %v", err)
	}

	if isSudo {
		command = strings.TrimPrefix(command, "sudo ")
	}

	if !b.isAuthorized(message.From.ID) {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Unauthorized access")
		b.api.Send(msg)
		return
	}

	// Execute the command
	output, err := b.executeCommand(command)
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Error: %v\n%s", err, output))
		b.api.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, output)
	b.api.Send(msg)
}

func loadConfig(configFile string) (Config, error) {
	file, err := os.Open(configFile)
	if err != nil {
		return Config{}, fmt.Errorf("failed to open config file: %v", err)
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return Config{}, fmt.Errorf("failed to parse config file: %v", err)
	}

	return config, nil
}

func main() {
  cmd := exec.Command("clear")
  cmd.Stdout = os.Stdout
  cmd.Run()

  printBanner()

	config, err := loadConfig("config.json")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	bot, err := NewBot(config, "bot_commands.log")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Bot started")

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := bot.api.GetUpdatesChan(updateConfig)
	updates.Clear()

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("Message from UserID: %d", update.Message.From.ID)

		switch {
		case update.Message.Command() == "start":
			bot.handleStart(update.Message)
		case update.Message.Command() == "getlogs":
			bot.handleGetLogs(update.Message)
		// case update.Message.Command() == "sudopassword":
		// 	bot.handleSudoPassword(update.Message)
	case update.Message.Command() == "listscripts":
		   bot.listScripts(update.Message)
		 case update.Message.Command() == "runscript":
		   bot.handleScript(update.Message)
		default:
			bot.handleCommand(update.Message)
		}
	}
}







































































































































































































































































































































































































































































































































































































































































































































































































































































































































































































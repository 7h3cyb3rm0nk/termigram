package bot

import (
	// "encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Config struct {
	BotToken     string   `json:"bot_token"`
	AllowedUsers []int64  `json:"allowed_users"`
	Scripts      []Script `json:"scripts"`
}

type Script struct {
	ID      string `json:"id"`
	Command string `json:"command"`
	Comment string `json:"comment"`
}

type Bot struct {
	Api            *tgbotapi.BotAPI
	AllowedUserIDs map[int64]bool   
	LogFile        string           
	WorkingDir     string           
	Config         Config 
	MessageChan    chan string
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
		Api:            api,
		AllowedUserIDs: allowedMap,
		LogFile:        logFile,
		WorkingDir:     os.Getenv("."),
		Config:         config,
		MessageChan:    make(chan string, 100),
	}, nil
}

func (b *Bot) isAuthorized(userID int64) bool {
	return b.AllowedUserIDs[userID]
	// return true
}


func (b *Bot) logCommand(userID int64, command string) error {
	f, err := os.OpenFile(b.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		b.MessageChan <- fmt.Sprintf("[ERROR] Failed to open log file: %v", err)
    return err
	}
	defer f.Close()

	logEntry := fmt.Sprintf("[%s] UserID: %d, Command: %s\n",
		time.Now().Format(time.RFC3339),
		userID,
		command)

	if _, err := f.WriteString(logEntry); err != nil {
    b.MessageChan <- fmt.Sprintf("[ERROR] Failed to write to log file: %v", err)
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
		b.WorkingDir, _ = os.Getwd() 
		return fmt.Sprintf("Changed directory to: %s", b.WorkingDir), nil
	} else {
		cmd = exec.Command("/bin/sh", "-c", command)
		cmd.Dir = b.WorkingDir 
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
	b.Api.Send(msg)
}

func (b *Bot) handleGetLogs(message *tgbotapi.Message) {

	if !b.isAuthorized(message.From.ID) {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Unauthorized access")
		b.Api.Send(msg)
		return
	}
	logs, err := os.ReadFile(b.LogFile)
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Failed to read logs")
		b.Api.Send(msg)
		return
	}

	response := "Recent command logs:\n" + string(logs)
	msg := tgbotapi.NewMessage(message.Chat.ID, response)
	b.Api.Send(msg)
}


func (b *Bot) listScripts(message *tgbotapi.Message){
	if !b.isAuthorized(message.From.ID) {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Unauthorized access")
		b.Api.Send(msg)
		return
	}

	response := "";
	for _, script := range b.Config.Scripts{
		response += script.ID +"   "+ script.Comment + "    " + script.Command
		response += "\n"

	}
	msg := tgbotapi.NewMessage(message.Chat.ID, response)
	b.Api.Send(msg)

}

func (b *Bot) handleScript(message *tgbotapi.Message) {
	if !b.isAuthorized(message.From.ID) {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Unauthorized access")
		b.Api.Send(msg)
		return
	}

	scriptID := strings.TrimPrefix(message.Text, "/runscript ") 
	var scriptToRun Script
	found := false
	for _, s := range b.Config.Scripts {
		if s.ID == scriptID {
			scriptToRun = s
			found = true
			break
		}
	}

	if !found {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Script not found")
		b.Api.Send(msg)
		return
	}

	// Execute the script
	Command := "./" + scriptToRun.Command
	output, err := b.executeCommand(Command)
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Error running script: %v\n%s", err, output))
		b.Api.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Script output:\n%s", output))
	b.Api.Send(msg)
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
		b.Api.Send(msg)
		return
	}

	// Execute the command
	output, err := b.executeCommand(command)
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Error: %v\n%s", err, output))
		b.Api.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, output)
	b.Api.Send(msg)
}

func (b *Bot) Start() {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := b.Api.GetUpdatesChan(updateConfig)
	updates.Clear()

	for update := range updates {
		if update.Message == nil {
			continue
		}

		// Send message to UI through channel
		b.MessageChan <- fmt.Sprintf("[%s] UserID: %d, Command: %s",
			time.Now().Format("15:04:05"),
			update.Message.From.ID,
			update.Message.Text,
		)

		switch {
		case update.Message.Command() == "start":
			b.handleStart(update.Message)
		case update.Message.Command() == "getlogs":
			b.handleGetLogs(update.Message)
		case update.Message.Command() == "listscripts":
			b.listScripts(update.Message)
		case update.Message.Command() == "runscript":
			b.handleScript(update.Message)
		default:
			b.handleCommand(update.Message)
		}
	}
}

package main

import (
    "encoding/json"
    "fmt"
    "log"
    "os"
    "os/exec"
    "strings"
    "time"
    "github.com/joho/godotenv"
    "sync"

    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)
var (
	sudoPasswd string
	mutex       sync.RWMutex // Add a mutex for safe concurrent access
)
type Config struct {
    BotToken     string   `json:"bot_token"`
    AllowedUsers []int64  `json:"allowed_users"`
}

type Bot struct {
    api            *tgbotapi.BotAPI
    allowedUserIDs map[int64]bool
    logFile        string
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
    }, nil
}

func (b *Bot) isAuthorized(userID int64) bool {
    return b.allowedUserIDs[userID]
    // return true
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

func (b *Bot) executeCommand(command string, sudo bool) (string, error) {
    var cmd *exec.Cmd

    if sudo {
        cmd = exec.Command("sudo", append([]string{"-n"}, strings.Fields(command)...)...)
    } else {
        cmd = exec.Command("/bin/sh", "-c", command)
    }

    output, err := cmd.CombinedOutput()
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
        "/sudopassword {password} - to pass the password\n" +
        "For regular commands, just type the command\n" +
        "For sudo commands, prefix with 'sudo '\n\n" +
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
    output, err := b.executeCommand(command, isSudo)
    if err != nil {
        msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Error: %v\n%s", err, output))
        b.api.Send(msg)
        return
    }

    msg := tgbotapi.NewMessage(message.Chat.ID, output)
    b.api.Send(msg)
}

func (b *Bot ) handleSudoPassword(message *tgbotapi.Message){
    // passwd := message.CommandArguments()[0]
    // mutex.Lock()
    // sudoPasswd = fmt.Sprintf("%s", passwd)
    // fmt.Println(passwd)
    // mutex.Unlock()
    // return true

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
    godotenv.Load(".env")

    // Load config from file
    config, err := loadConfig("config.json")
    if err != nil {
        log.Fatalf("Error loading config: %v", err)
    }

    // Initialize the bot with your token and allowed user IDs from the config
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

        // Log or show the user ID without restriction
        log.Printf("Message from UserID: %d", update.Message.From.ID)

        // Handle different types of commands
        switch {
        case update.Message.Command() == "start":
            bot.handleStart(update.Message)
        case update.Message.Command() == "getlogs":
            bot.handleGetLogs(update.Message)
        case update.Message.Command() == "sudopassword":
            bot.handleSudoPassword(update.Message)
        default:
            bot.handleCommand(update.Message)
        }
    }
}



























































































































































































































































































































































































































































































































































































































































































































































































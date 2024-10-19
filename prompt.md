Project Path: /home/ramees/progs/go/termigram

Source Tree:

```
termigram
├── go.mod
├── bot_commands.log
├── go.sum
├── server
└── bot
    ├── bot.go
    └── main.go

```

`/home/ramees/progs/go/termigram/go.mod`:

```````mod
module github.com/7h3cyb3rm0nk/termigram

go 1.22.5

require (
	github.com/go-telegram-bot-api/telegram-bot-api/v5 v5.5.1 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	golang.org/x/crypto v0.28.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

```````

`/home/ramees/progs/go/termigram/bot_commands.log`:

```````log
[2024-10-19T22:36:36+05:30] UserID: 1257658173, Command: ls
[2024-10-19T22:36:41+05:30] UserID: 1257658173, Command: getId
[2024-10-19T22:36:47+05:30] UserID: 1257658173, Command: /getId
[2024-10-19T22:43:52+05:30] UserID: 1257658173, Command: ls
[2024-10-19T22:44:05+05:30] UserID: 1257658173, Command: sudo
[2024-10-19T22:44:13+05:30] UserID: 1257658173, Command: su
[2024-10-19T22:45:19+05:30] UserID: 7991366741, Command: Helki
[2024-10-19T22:45:30+05:30] UserID: 7991366741, Command: Ls
[2024-10-19T22:45:34+05:30] UserID: 7991366741, Command: ls
[2024-10-19T23:15:11+05:30] UserID: 1257658173, Command: /sudoPassword
[2024-10-19T23:15:34+05:30] UserID: 1257658173, Command: su
[2024-10-19T23:19:25+05:30] UserID: 1257658173, Command: su
[2024-10-19T23:21:18+05:30] UserID: 1257658173, Command: su
[2024-10-19T23:27:49+05:30] UserID: 1257658173, Command: /sudopasswd k3v1n
[2024-10-19T23:27:58+05:30] UserID: 1257658173, Command: /sudopasswd
[2024-10-19T23:28:06+05:30] UserID: 1257658173, Command: /sudoPasswd
[2024-10-19T23:32:10+05:30] UserID: 7991366741, Command: 🖖
[2024-10-19T23:32:43+05:30] UserID: 7302237680, Command: ls
[2024-10-19T23:32:53+05:30] UserID: 7991366741, Command: ls
[2024-10-19T23:32:58+05:30] UserID: 7991366741, Command: ls
[2024-10-19T23:42:47+05:30] UserID: 7991366741, Command: ls
[2024-10-19T23:42:55+05:30] UserID: 7991366741, Command: sudo
[2024-10-19T23:44:49+05:30] UserID: 7991366741, Command: Hi
[2024-10-19T23:44:58+05:30] UserID: 7991366741, Command: ls
[2024-10-19T23:49:47+05:30] UserID: 7991366741, Command: sudo su
[2024-10-19T23:49:54+05:30] UserID: 7991366741, Command: whoami
[2024-10-20T01:58:04+05:30] UserID: 7991366741, Command: whoami
[2024-10-20T01:58:11+05:30] UserID: 7991366741, Command: mkdir test
[2024-10-20T01:58:13+05:30] UserID: 7991366741, Command: ls
[2024-10-20T01:58:22+05:30] UserID: 7991366741, Command: rm -rf test
[2024-10-20T01:58:25+05:30] UserID: 7991366741, Command: ls
[2024-10-20T01:58:33+05:30] UserID: 7991366741, Command: cd ..
[2024-10-20T01:58:36+05:30] UserID: 7991366741, Command: pwd
[2024-10-20T01:58:44+05:30] UserID: 7991366741, Command: cs ../
[2024-10-20T01:58:51+05:30] UserID: 7991366741, Command: cd ../
[2024-10-20T01:58:54+05:30] UserID: 7991366741, Command: pwd
[2024-10-20T01:59:26+05:30] UserID: 7991366741, Command: cd /home/ramees/
[2024-10-20T01:59:31+05:30] UserID: 7991366741, Command: ls
[2024-10-20T02:00:07+05:30] UserID: 7991366741, Command: ls

```````

`/home/ramees/progs/go/termigram/go.sum`:

```````sum
github.com/go-telegram-bot-api/telegram-bot-api/v5 v5.5.1 h1:wG8n/XJQ07TmjbITcGiUaOtXxdrINDz1b0J1w0SzqDc=
github.com/go-telegram-bot-api/telegram-bot-api/v5 v5.5.1/go.mod h1:A2S0CWkNylc2phvKXWBBdD3K0iGnDBGbzRpISP2zBl8=
github.com/joho/godotenv v1.5.1 h1:7eLL/+HRGLY0ldzfGMeQkb7vMd0as4CfYvUVzLqw0N0=
github.com/joho/godotenv v1.5.1/go.mod h1:f4LDr5Voq0i2e/R5DDNOoa2zzDfwtkZa6DnEwAbqwq4=
golang.org/x/crypto v0.28.0 h1:GBDwsMXVQi34v5CCYUm2jkJvu4cbtru2U4TN2PSyQnw=
golang.org/x/crypto v0.28.0/go.mod h1:rmgy+3RHxRZMyY0jjAJShp2zgEdOqj2AO7U0pYmeQ7U=
golang.org/x/sys v0.26.0 h1:KHjCJyddX0LoSTb3J+vWpupP9p0oznkqVk/IfjymZbo=
golang.org/x/sys v0.26.0/go.mod h1:/VUhepiaJMQUp4+oa/7Zr1D23ma6VTLIYjOOTFZPUcA=
gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405/go.mod h1:Co6ibVJAznAaIkqp8huTwlJQCZ016jof/cbN4VW5Yz0=
gopkg.in/yaml.v2 v2.4.0 h1:D8xgwECY7CYvx+Y2n4sBz93Jn9JRvxdiyyo8CTfuKaY=
gopkg.in/yaml.v2 v2.4.0/go.mod h1:RDklbk79AGWmwhnvt/jBztapEOGDOx6ZbXqjP6csGnQ=

```````

`/home/ramees/progs/go/termigram/bot/bot.go`:

```````go
package main

import (
    "bytes"
    "crypto/rsa"
    "crypto/sha256"
    "crypto/tls"
    "crypto/x509"
    "encoding/base64"
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
    "os/exec"
    "path/filepath"
    "strings"
    "sync"
    "time"

    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
    "github.com/joho/godotenv"
    "golang.org/x/crypto/ssh"
    "gopkg.in/yaml.v2"
)

// Role defines permissions and access levels
type Role struct {
    Name            string
    Level           int
    AllowedCommands []string
    AllowedScripts  []string
    Permissions     []string
    MaxTimeout      time.Duration
    RateLimit       RateLimit
}

// RateLimit defines rate limiting parameters
type RateLimit struct {
    RequestsPerMinute int
    BurstSize        int
    CooldownPeriod   time.Duration
}

// User represents a system user with roles and permissions
type User struct {
    ID          int64
    Roles       []string
    Auth0ID     string
    MFAEnabled  bool
    IPWhitelist []string
    PublicKey   string
    LastAccess  time.Time
    Status      string
}

// Script represents a managed script
type Script struct {
    ID          string
    Name        string
    Description string
    Content     string
    Hash        string
    Version     int
    Author      string
    Created     time.Time
    Modified    time.Time
    RequiredRole string
    ApprovalFlow []string
    Timeout     time.Duration
    Environment map[string]string
}

// RemoteServer represents a managed remote server
type RemoteServer struct {
    ID          string
    Name        string
    Host        string
    Port        int
    User        string
    KeyPath     string
    AllowedIPs  []string
    Environment map[string]string
    Permissions map[string][]string
}

// RBACManager handles role-based access control
type RBACManager struct {
    mu           sync.RWMutex
    roles        map[string]Role
    users        map[int64]User
    permissions  map[string][]string
    auditLogger  *AuditLogger
}

// ScriptManager handles script management and execution
type ScriptManager struct {
    mu            sync.RWMutex
    scripts       map[string]Script
    versions      map[string][]Script
    approvals     map[string][]Approval
    scriptDir     string
    backupDir     string
}

// RemoteExecutor handles secure remote command execution
type RemoteExecutor struct {
    mu             sync.RWMutex
    servers        map[string]RemoteServer
    connections    map[string]*ssh.Client
    keyring        ssh.PublicKeys
    timeout        time.Duration
}

// Approval represents a script approval workflow
type Approval struct {
    ScriptID    string
    Version     int
    ApproverID  int64
    Status      string
    Comment     string
    Timestamp   time.Time
}

// Enhanced Config structure
type Config struct {
    // ... previous config fields ...
    RBAC          RBACConfig
    Scripts       ScriptConfig
    RemoteServers []RemoteServer
}

// RBACConfig holds RBAC-specific configuration
type RBACConfig struct {
    Roles           map[string]Role
    DefaultRole     string
    AdminUsers      []int64
    RequireMFA      bool
    ApprovalFlows   map[string][]string
}

// ScriptConfig holds script management configuration
type ScriptConfig struct {
    ScriptDir     string
    BackupDir     string
    MaxScriptSize int64
    AutoApprove   bool
    RequireReview bool
}

// Enhanced Bot structure
type Bot struct {
    // ... previous bot fields ...
    rbac     *RBACManager
    scripts  *ScriptManager
    remote   *RemoteExecutor
}

// NewRBACManager creates a new RBAC manager
func NewRBACManager(config RBACConfig, logger *AuditLogger) *RBACManager {
    return &RBACManager{
        roles:       config.Roles,
        users:      make(map[int64]User),
        permissions: make(map[string][]string),
        auditLogger: logger,
    }
}

// AddUser adds a new user with roles
func (rm *RBACManager) AddUser(user User) error {
    rm.mu.Lock()
    defer rm.mu.Unlock()

    // Validate user roles
    for _, role := range user.Roles {
        if _, exists := rm.roles[role]; !exists {
            return fmt.Errorf("invalid role: %s", role)
        }
    }

    rm.users[user.ID] = user
    rm.auditLogger.LogEvent(AuditLog{
        UserID:    user.ID,
        Event:     "USER_ADDED",
        Status:    "SUCCESS",
        Timestamp: time.Now(),
    })

    return nil
}

// CheckPermission checks if a user has required permissions
func (rm *RBACManager) CheckPermission(userID int64, permission string) bool {
    rm.mu.RLock()
    defer rm.mu.RUnlock()

    user, exists := rm.users[userID]
    if !exists {
        return false
    }

    for _, roleName := range user.Roles {
        role, exists := rm.roles[roleName]
        if !exists {
            continue
        }

        for _, perm := range role.Permissions {
            if perm == permission || perm == "*" {
                return true
            }
        }
    }

    return false
}

// NewScriptManager creates a new script manager
func NewScriptManager(config ScriptConfig) (*ScriptManager, error) {
    // Create script and backup directories
    for _, dir := range []string{config.ScriptDir, config.BackupDir} {
        if err := os.MkdirAll(dir, 0750); err != nil {
            return nil, err
        }
    }

    return &ScriptManager{
        scripts:   make(map[string]Script),
        versions:  make(map[string][]Script),
        approvals: make(map[string][]Approval),
        scriptDir: config.ScriptDir,
        backupDir: config.BackupDir,
    }, nil
}

// AddScript adds a new script with version control
func (sm *ScriptManager) AddScript(script Script) error {
    sm.mu.Lock()
    defer sm.mu.Unlock()

    // Generate script hash
    hash := sha256.Sum256([]byte(script.Content))
    script.Hash = base64.StdEncoding.EncodeToString(hash[:])

    // Check if script already exists
    if existing, exists := sm.scripts[script.ID]; exists {
        script.Version = existing.Version + 1
    } else {
        script.Version = 1
    }

    // Save script file
    filename := filepath.Join(sm.scriptDir, fmt.Sprintf("%s_v%d.sh", script.ID, script.Version))
    if err := os.WriteFile(filename, []byte(script.Content), 0640); err != nil {
        return err
    }

    // Update script records
    sm.scripts[script.ID] = script
    sm.versions[script.ID] = append(sm.versions[script.ID], script)

    return nil
}

// NewRemoteExecutor creates a new remote executor
func NewRemoteExecutor(servers []RemoteServer, timeout time.Duration) (*RemoteExecutor, error) {
    executor := &RemoteExecutor{
        servers:     make(map[string]RemoteServer),
        connections: make(map[string]*ssh.Client),
        timeout:    timeout,
    }

    // Load SSH key
    key, err := os.ReadFile(os.Getenv("SSH_PRIVATE_KEY"))
    if err != nil {
        return nil, fmt.Errorf("failed to load SSH key: %v", err)
    }

    signer, err := ssh.ParsePrivateKey(key)
    if err != nil {
        return nil, fmt.Errorf("failed to parse SSH key: %v", err)
    }

    executor.keyring = signer

    // Initialize servers
    for _, server := range servers {
        executor.servers[server.ID] = server
    }

    return executor, nil
}

// ExecuteRemoteCommand executes a command on a remote server
func (re *RemoteExecutor) ExecuteRemoteCommand(serverID, command string, user User) (string, error) {
    re.mu.Lock()
    defer re.mu.Unlock()

    server, exists := re.servers[serverID]
    if !exists {
        return "", fmt.Errorf("server not found: %s", serverID)
    }

    // Check IP whitelist
    if len(server.AllowedIPs) > 0 {
        ipAllowed := false
        for _, ip := range user.IPWhitelist {
            for _, allowedIP := range server.AllowedIPs {
                if ip == allowedIP {
                    ipAllowed = true
                    break
                }
            }
        }
        if !ipAllowed {
            return "", fmt.Errorf("IP not allowed")
        }
    }

    // Get or create SSH connection
    client, err := re.getSSHConnection(server)
    if err != nil {
        return "", err
    }

    // Create session
    session, err := client.NewSession()
    if err != nil {
        return "", err
    }
    defer session.Close()

    // Set up output buffers
    var stdout, stderr bytes.Buffer
    session.Stdout = &stdout
    session.Stderr = &stderr

    // Set environment variables
    for k, v := range server.Environment {
        if err := session.Setenv(k, v); err != nil {
            return "", err
        }
    }

    // Execute command with timeout
    done := make(chan error, 1)
    go func() {
        done <- session.Run(command)
    }()

    select {
    case err := <-done:
        if err != nil {
            return "", fmt.Errorf("execution error: %v\nstderr: %s", err, stderr.String())
        }
    case <-time.After(re.timeout):
        session.Signal(ssh.SIGTERM)
        return "", fmt.Errorf("command timed out")
    }

    return stdout.String(), nil
}

// Example of enhanced command handling
func (b *Bot) handleCommand(message *tgbotapi.Message) {
    userID := message.From.ID
    command := strings.TrimPrefix(message.Text, b.config.CommandPrefix)

    // Check user permissions
    if !b.rbac.CheckPermission(userID, "execute_command") {
        b.sendMessage(message.Chat.ID, "Permission denied")
        return
    }

    // Parse command
    parts := strings.Fields(command)
    if len(parts) < 2 {
        b.sendMessage(message.Chat.ID, "Invalid command format. Use: /cmd <server> <command>")
        return
    }

    serverID := parts[0]
    cmdStr := strings.Join(parts[1:], " ")

    // Get user
    user, err := b.rbac.GetUser(userID)
    if err != nil {
        b.sendMessage(message.Chat.ID, "User not found")
        return
    }

    // Execute command
    output, err := b.remote.ExecuteRemoteCommand(serverID, cmdStr, user)
    if err != nil {
        b.sendMessage(message.Chat.ID, fmt.Sprintf("Error: %v", err))
        return
    }

    b.sendMessage(message.Chat.ID, fmt.Sprintf("Output:\n%s", output))
}

// Example main function with enhanced configuration
func main() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    // Load RBAC configuration
    rbacConfig := RBACConfig{
        Roles: map[string]Role{
            "admin": {
                Name:            "admin",
                Level:           100,
                AllowedCommands: []string{"*"},
                AllowedScripts:  []string{"*"},
                Permissions:     []string{"*"},
                MaxTimeout:      time.Minute * 30,
                RateLimit: RateLimit{
                    RequestsPerMinute: 60,
                    BurstSize:        10,
                    CooldownPeriod:   time.Minute,
                },
            },
            "operator": {
                Name:            "operator",
                Level:           50,
                AllowedCommands: []string{"status", "restart", "logs"},
                AllowedScripts:  []string{"backup", "deploy"},
                Permissions:     []string{"execute_command", "view_logs"},
                MaxTimeout:      time.Minute * 15,
                RateLimit: RateLimit{
                    RequestsPerMinute: 30,
                    BurstSize:        5,
                    CooldownPeriod:   time.Minute * 2,
                },
            },
            "viewer": {
                Name:            "viewer",
                Level:           10,
                AllowedCommands: []string{"status", "logs"},
                AllowedScripts:  []string{"report"},
                Permissions:     []string{"view_logs"},
                MaxTimeout:      time.Minute * 5,
                RateLimit: RateLimit{
                    RequestsPerMinute: 10,
                    BurstSize:        3,
                    CooldownPeriod:   time.Minute * 5,
                },
            },
        },
        DefaultRole: "viewer",
        AdminUsers:  []int64{/* Add admin user IDs */},
        RequireMFA:  true,
    }

    // Initialize remote servers
    remoteServers := []RemoteServer{
        {
            ID:         "prod-1",
            Name:       "Production Server 1",
            Host:       os.Getenv("PROD_SERVER_1_HOST"),
            Port:       22,
            User:       "deploy",
            KeyPath:    "/path/to/ssh/key",
            AllowedIPs: []string{"10.0.0.0/24"},
            Environment: map[string]string{
                "ENV": "production",
            },
            Permissions: map[string][]string{
                "admin":    {"*"},
                "operator": {"restart", "deploy"},
                "viewer":   {"status"},
            },
        },
        // Add more servers as needed
    }

    // Start the bot with enhanced configuration
    config := &Config{
        // ... previous config fields ...
        RBAC:          rbacConfig,
        Scripts: ScriptConfig{
            ScriptDir:     "/opt/bot/scripts",
            BackupDir:     "/opt/bot/backups",
            MaxScriptSize: 1024 * 1024, // 1MB
            AutoApprove:   false,
            RequireReview: true,
        },
        RemoteServers: remoteServers,
    }

    bot, err := NewBot(config)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Secure RBAC bot started")
    if err := bot.Start(); err != nil {
        log.Fatal(err)
    }
}

```````

`/home/ramees/progs/go/termigram/bot/main.go`:

```````go
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
    "sync"

    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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
            // bot.handleSudoPassword(update.Message)
        default:
            bot.handleCommand(update.Message)
        }
    }
}



























































































































































































































































































































































































































































































































































































































































































































































































```````
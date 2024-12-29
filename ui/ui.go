package ui

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/7h3cyb3rm0nk/termigram/bot"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"golang.org/x/term"
)

type StoredCredential struct {
	UserID   int64  `json:"user_id"`
	BotToken string `json:"bot_token"`
}

type TermUI struct {
	bot           *bot.Bot
	startTime     time.Time
	header        *widgets.Paragraph
	logs          *widgets.List
	stats         *widgets.Paragraph
	input         *widgets.Paragraph
	credentials   *widgets.List
	setup         bool
	storedCreds   []StoredCredential
	credsFilePath string
}

func NewTermUI(bot *bot.Bot) *TermUI {
	return &TermUI{
		startTime:     time.Now(),
		setup:         false,
		credsFilePath: "credentials.json",
	}
}

func (t *TermUI) setupWidgets() {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		width = 80
	}

	t.header = widgets.NewParagraph()
	t.header.Title = "Termigram Bot"
	t.header.SetRect(0, 0, width, 3)
	t.header.BorderStyle.Fg = ui.ColorCyan

	t.credentials = widgets.NewList()
	t.credentials.Title = "Stored Credentials"
	t.credentials.SetRect(0, 3, width, 10)
	t.credentials.BorderStyle.Fg = ui.ColorBlue

	t.input = widgets.NewParagraph()
	t.input.Title = "Configuration"
	t.input.SetRect(0, 10, width, 15)
	t.input.BorderStyle.Fg = ui.ColorMagenta

	t.logs = widgets.NewList()
	t.logs.Title = "Recent Commands"
	t.logs.SetRect(0, 15, width, 22)
	t.logs.BorderStyle.Fg = ui.ColorYellow

	t.stats = widgets.NewParagraph()
	t.stats.Title = "Statistics"
	t.stats.SetRect(0, 22, width, 27)
	t.stats.BorderStyle.Fg = ui.ColorGreen
}

func (t *TermUI) getInput(prompt string) (string, error) {
	if t.input == nil {
		return "", fmt.Errorf("input widget not initialized")
	}

	inputText := ""
	t.input.Text = prompt + "\n" + inputText
	t.updateUI()

	for {
		e := <-ui.PollEvents()
		switch e.ID {
		case "<C-c>":
			return "", fmt.Errorf("cancelled")
		case "<Backspace>":
			if len(inputText) > 0 {
				inputText = inputText[:len(inputText)-1]
			}
		case "<Enter>":
			if len(inputText) > 0 {
				return inputText, nil
			}
		case "<Space>":
			inputText += " "
		default:
			if len(e.ID) == 1 {
				inputText += e.ID
			}
		}
		t.input.Text = fmt.Sprintf("%s\n%s", prompt, inputText)
		t.updateUI()
	}
}

func (t *TermUI) updateUI() {
	if t.credentials != nil && t.input != nil && t.logs != nil && t.stats != nil {
		ui.Render(t.header, t.credentials, t.input, t.logs, t.stats)
	}
}

func maskToken(token string) string {
	if len(token) <= 8 {
		return strings.Repeat("*", len(token))
	}
	return token[:4] + strings.Repeat("*", len(token)-8) + token[len(token)-4:]
}

func (t *TermUI) loadCredentials() error {
	data, err := os.ReadFile(t.credsFilePath)
	if os.IsNotExist(err) {
		t.storedCreds = []StoredCredential{}
		return nil
	}
	if err != nil {
		return fmt.Errorf("failed to read credentials: %v", err)
	}
	if err := json.Unmarshal(data, &t.storedCreds); err != nil {
		return fmt.Errorf("failed to parse credentials: %v", err)
	}
	return nil
}

func (t *TermUI) saveCredentials() error {
	data, err := json.MarshalIndent(t.storedCreds, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal credentials: %v", err)
	}
	return os.WriteFile(t.credsFilePath, data, 0600)
}

func (t *TermUI) updateCredentialsDisplay() {
	t.credentials.Rows = make([]string, 0, len(t.storedCreds)+2)
	for i, cred := range t.storedCreds {
		t.credentials.Rows = append(t.credentials.Rows,
			fmt.Sprintf("[%d] UserID: %d | Token: %s", i+1, cred.UserID, maskToken(cred.BotToken)))
	}
	t.credentials.Rows = append(t.credentials.Rows,
		"[N] Add New Credentials",
		"[D] Delete Credentials")
}

func (t *TermUI) configureBot() error {
	for {
		if err := t.loadCredentials(); err != nil {
			return err
		}

		t.updateCredentialsDisplay()
		t.updateUI()

		choice, err := t.getInput("Enter index to use existing, 'N' for new, or 'D' to delete:")
		if err != nil {
			return err
		}

		choice = strings.TrimSpace(strings.ToUpper(choice))

		switch choice {
		case "N":
			if err := t.handleNewCredentials(); err != nil {
				t.showError(err.Error())
				continue
			}
			return nil
		case "D":
			if err := t.handleDeleteCredentials(); err != nil {
				t.showError(err.Error())
				continue
			}
		default:
			if err := t.handleExistingCredentials(choice); err != nil {
				t.showError(err.Error())
				continue
			}
			return nil
		}
	}
}

func (t *TermUI) handleNewCredentials() error {
	token, err := t.getInput("Enter new Telegram Bot API Token:")
	if err != nil {
		return err
	}

	userIDStr, err := t.getInput("Enter your User ID:")
	if err != nil {
		return err
	}

	var userID int64
	if _, err := fmt.Sscanf(userIDStr, "%d", &userID); err != nil {
		return fmt.Errorf("invalid user ID format")
	}

	t.storedCreds = append(t.storedCreds, StoredCredential{
		UserID:   userID,
		BotToken: token,
	})

	if err := t.saveCredentials(); err != nil {
		return err
	}

	return t.initializeBot(token, userID)
}

func (t *TermUI) handleDeleteCredentials() error {
	if len(t.storedCreds) == 0 {
		return fmt.Errorf("no credentials to delete")
	}

	indexStr, err := t.getInput("Enter index number to delete:")
	if err != nil {
		return err
	}

	var index int
	if _, err := fmt.Sscanf(indexStr, "%d", &index); err != nil || index < 1 || index > len(t.storedCreds) {
		return fmt.Errorf("invalid index")
	}

	t.storedCreds = append(t.storedCreds[:index-1], t.storedCreds[index:]...)
	return t.saveCredentials()
}

func (t *TermUI) handleExistingCredentials(choice string) error {
	var index int
	if _, err := fmt.Sscanf(choice, "%d", &index); err != nil || index < 1 || index > len(t.storedCreds) {
		return fmt.Errorf("invalid selection")
	}

	cred := t.storedCreds[index-1]
	return t.initializeBot(cred.BotToken, cred.UserID)
}

func (t *TermUI) initializeBot(token string, userID int64) error {
	config := bot.Config{
		BotToken:     token,
		AllowedUsers: []int64{userID},
		Scripts:      []bot.Script{},
	}

	newBot, err := bot.NewBot(config, "bot_commands.log")
	if err != nil {
		return fmt.Errorf("failed to initialize bot: %v", err)
	}

	t.bot = newBot
	t.setup = true
	return nil
}

func (t *TermUI) showError(msg string) {
	t.input.Text = fmt.Sprintf("Error: %s\nPress any key to continue...", msg)
	t.updateUI()
	<-ui.PollEvents()
}

func (t *TermUI) Start() error {
	if err := ui.Init(); err != nil {
		return fmt.Errorf("failed to initialize UI: %v", err)
	}
	defer ui.Close()

	t.setupWidgets()
	t.updateUI()

	if !t.setup {
		if err := t.configureBot(); err != nil {
			return err
		}
	}

	go t.bot.Start()

	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Second).C

	for {
		select {
		case e := <-uiEvents:
			if e.ID == "q" || e.ID == "<C-c>" {
				return nil
			}
		case msg := <-t.bot.MessageChan:
			t.updateLogs(msg)
		case <-ticker:
			t.updateStats()
		}
	}
}

func (t *TermUI) updateLogs(msg string) {
	t.logs.Rows = append([]string{msg}, t.logs.Rows...)
	if len(t.logs.Rows) > 10 {
		t.logs.Rows = t.logs.Rows[:10]
	}
	t.updateUI()
}

func (t *TermUI) updateStats() {
	t.stats.Text = fmt.Sprintf(
		"Uptime: %v\nAuthorized Users: %d\nScripts Available: %d",
		time.Since(t.startTime).Round(time.Second),
		len(t.bot.AllowedUserIDs),
		len(t.bot.Config.Scripts),
	)
	t.updateUI()
}

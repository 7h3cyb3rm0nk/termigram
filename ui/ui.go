package ui

import (
	"fmt"
	// "log"
	"time"

	"github.com/7h3cyb3rm0nk/termigram/bot"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	 "golang.org/x/term"
	 "os"
)

type TermUI struct {
	bot       *bot.Bot
	startTime time.Time
	header    *widgets.Paragraph
	logs      *widgets.List
	stats     *widgets.Paragraph
}

func NewTermUI(bot *bot.Bot) *TermUI {
	return &TermUI{
		bot:       bot,
		startTime: time.Now(),
	}
}

func (t *TermUI) setupWidgets() {

	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil{
		width=80
	}
	t.header = widgets.NewParagraph()
	t.header.Title = "Termigram Bot"
	t.header.Text = fmt.Sprintf("Bot Username: @%s | Status: Running", t.bot.Api.Self.UserName)
	t.header.SetRect(0, 0, width, 3)
	t.header.BorderStyle.Fg = ui.ColorCyan

	t.logs = widgets.NewList()
	t.logs.Title = "Recent Commands"
	t.logs.SetRect(0, 3, width, 15)
	t.logs.BorderStyle.Fg = ui.ColorYellow

	t.stats = widgets.NewParagraph()
	t.stats.Title = "Statistics"
	t.stats.SetRect(0, 15, width, 20)
	t.stats.BorderStyle.Fg = ui.ColorGreen
}

func (t *TermUI) updateUI() {
	ui.Render(t.header, t.logs, t.stats)
}

func (t *TermUI) Start() error {
	if err := ui.Init(); err != nil {
		return fmt.Errorf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	t.setupWidgets()
	t.updateUI()

	// Handle events
	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Second).C

	go t.bot.Start() 

	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return nil
			}
		case msg := <-t.bot.MessageChan:
			// Update logs
			t.logs.Rows = append([]string{msg}, t.logs.Rows...)
			if len(t.logs.Rows) > 10 {
				t.logs.Rows = t.logs.Rows[:10]
			}
			t.updateUI()
		case <-ticker:
			// Update stats
			t.stats.Text = fmt.Sprintf(
				"Uptime: %v\nAuthorized Users: %d\nScripts Available: %d",
				time.Since(t.startTime).Round(time.Second),
				len(t.bot.AllowedUserIDs),
				len(t.bot.Config.Scripts),
			)
			t.updateUI()
		}
	}
}


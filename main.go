package main


import (
	"log"
   "fmt"
   "encoding/json"
   "os"
	"github.com/7h3cyb3rm0nk/termigram/bot"
	"github.com/7h3cyb3rm0nk/termigram/ui"
)




func loadConfig(configFile string) (bot.Config, error) {
	file, err := os.Open(configFile)
	if err != nil {
		return bot.Config{}, fmt.Errorf("failed to open config file: %v", err)
	}
	defer file.Close()

	var config bot.Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return bot.Config{}, fmt.Errorf("failed to parse config file: %v", err)
	}

	return config, nil
}

func main() {
	// Load configuration
	config, err := loadConfig("config.json")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Create new bot instance
	bot, err := bot.NewBot(config, "bot_commands.log")
	if err != nil {
		log.Fatal(err)
	}

	// Create and start UI
	termUI := ui.NewTermUI(bot)
	if err := termUI.Start(); err != nil {
		log.Fatal(err)
	}
}






































































































































































































































































































































































































































































































































































































































































































































































































































































































































































































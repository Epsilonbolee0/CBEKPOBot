package main

import (
	"encoding/xml"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"io/ioutil"
	"log"
	"os"
)

const configPath = "./docker_compose/bot/src/config.xml"

func loadID() string {
	var botID string

	xmlFile, err := os.Open(configPath)
	if err != nil {
		panic(err)
	}
	defer xmlFile.Close()

	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		panic(err)
	}

	xml.Unmarshal(byteValue, &botID)
	return botID
}

func main() {
	id := loadID()
	bot, err := tgbotapi.NewBotAPI(id)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil || !update.Message.IsCommand() { // ignore any non-Message Updates
			continue
		}

		//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		switch update.Message.Command() {
		case "holiday":
			msg.Text = GetHolidayCommand()
		case "start":
			msg.ParseMode = "html"
			msg.Text = GetStartTextCommand()
		default:
			msg.ParseMode = "html"
			msg.Text = GetDefaultTextCommand()
		}
		bot.Send(msg)
	}
}

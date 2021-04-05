package bot

import (
	"encoding/xml"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"io/ioutil"
	"log"
	"os"
)

const (
	configPath    = "./docker_compose/src/bot/config.xml"
	updateTimeout = 60
)

type BotInterface interface {
	Run()
}

type Bot struct {
	telegramID string
	API        *tgbotapi.BotAPI
	CMD        Commands
}

func NewBot() *Bot {
	newBot := Bot{}
	newBot.loadID()
	newBot.setAPI()
	newBot.setCmd()

	return &newBot
}

func (bot *Bot) loadID() {
	var telegramID string

	xmlFile, err := os.Open(configPath)
	if err != nil {
		panic(err)
	}
	defer xmlFile.Close()

	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		panic(err)
	}

	xml.Unmarshal(byteValue, &telegramID)
	bot.telegramID = telegramID
}

func (bot *Bot) setAPI() {
	api, err := tgbotapi.NewBotAPI(bot.telegramID)
	if err != nil {
		panic(err)
	}

	api.Debug = false
	bot.API = api
}

func (bot *Bot) setCmd() {
	bot.CMD = InitCommands()
}

func (bot *Bot) Run() {
	bot.greet()
	bot.monitorMessageCycle()
}

func (bot *Bot) greet() {
	api := bot.API
	log.Printf("Авторизован под аккаунтом %s!", api.Self.UserName)
}

func (bot *Bot) monitorMessageCycle() {
	updates := bot.getUpdateChannel()

	for update := range updates {
		bot.handleUpdate(&update)
	}
}

func (bot *Bot) getUpdateChannel() tgbotapi.UpdatesChannel {
	api := bot.API

	update := tgbotapi.NewUpdate(0)
	update.Timeout = updateTimeout

	updateChannel, err := api.GetUpdatesChan(update)
	if err != nil {
		panic(err)
	}

	return updateChannel
}

func (bot *Bot) handleUpdate(update *tgbotapi.Update) {
	switch {
	case bot.updateIsCommand(update):
		bot.handleCommand(update)
		bot.logCommand(update)
	}
}

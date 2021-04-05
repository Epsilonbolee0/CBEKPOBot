package bot

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"log"
)

func (bot *Bot) logCommand(update *tgbotapi.Update) {
	log.Printf("%s запросил команду %s!\n", update.Message.From, update.Message.Text)
}

func (bot *Bot) updateIsCommand(update *tgbotapi.Update) bool {
	return update.Message != nil && update.Message.IsCommand()
}

func (bot *Bot) handleCommand(update *tgbotapi.Update) {
	callback := bot.getCommandCallback(update)
	msg := bot.newMessage(update, callback)
	bot.sendMessage(msg)
}

func (bot *Bot) getCommandCallback(update *tgbotapi.Update) CommandCallback {
	command := update.Message.Command()
	cmd := bot.CMD
	return cmd.GetCallback(command)
}

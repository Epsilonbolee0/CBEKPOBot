package bot

import tgbotapi "github.com/Syfaro/telegram-bot-api"

func (bot *Bot) newMessage(update *tgbotapi.Update, callback CommandCallback) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.ParseMode = "html"
	msg.Text = callback()

	return msg
}

func (bot *Bot) sendMessage(message tgbotapi.MessageConfig) {
	api := bot.API
	api.Send(message)
}

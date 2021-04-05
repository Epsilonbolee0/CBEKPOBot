package bot

import (
	fetcher "../fetchers/holiday_fetcher"
)

type CommandCallback func() string
type Commands map[string]CommandCallback

func InitCommands() Commands {
	cmd := Commands{}

	cmd["start"] = StartCommand
	cmd["holiday"] = HolidayCommand
	cmd["default"] = DefaultCommand
	cmd["info"] = InfoCommand

	return cmd
}

func (cmds Commands) GetCallback(goal string) CommandCallback {
	for command, callback := range cmds {
		if command == goal {
			return callback
		}
	}

	return cmds["default"]
}

func StartCommand() string {
	return "Приветик, мне тут телепонстер настроили.\n " +
		"Тётя Люба рассказала про эту <i>телепрограмму</i> вашу!\n " +
		"Буду теперь тебе открыточки слать :,) :0 :-))\n " +
		"<b>(Попроси /holiday! - всегда удружу!)</b>\n"
}

func InfoCommand() string {
	return "<b>/start</b> - Поздороваться" +
		"<b>/holiday</b> - Прислать случайный праздник сегодня"
}

func HolidayCommand() string {
	fetcher := fetcher.GetHolidayFetcher()
	return fetcher.GetRandomHoliday()
}

func DefaultCommand() string {
	return "Прости, не понимаю. Попроси <b>/holiday</b> !"
}

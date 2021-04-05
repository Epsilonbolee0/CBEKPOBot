package main

import (
	fetcher "./fetchers/holiday_fetcher"
)

func GetHolidayCommand() string {
	fetcher := fetcher.GetHolidayFetcher()
	return fetcher.GetRandomHoliday()
}

func GetDefaultTextCommand() string {
	return "Прости, не понимаю. Попроси <b>/holiday</b> !"
}

func GetStartTextCommand() string {
	return "Приветик, мне тут телепонстер настроили.\n " +
		"Тётя Люба рассказала про эту <i>телепрограмму</i> вашу!\n " +
		"Буду теперь тебе открыточки слать :,) :0 :-))\n " +
		"<b>(Попроси /holiday! - всегда удружу!)</b>\n"
}

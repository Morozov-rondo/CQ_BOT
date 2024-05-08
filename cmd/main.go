package main

import (
	CQ_bot "CQ-bot/config"
	"CQ-bot/internial/service/sheets"
	"CQ-bot/internial/service/tg_bot"
	"CQ-bot/models"
)

// Точка запуска приложения
func main() {

	CQ_bot.NewConfig()

	ChanTeam := make(chan models.TeamData)

	go sheets.ParsinNewTeam(ChanTeam)
	tg_bot.TgBot(ChanTeam)

}

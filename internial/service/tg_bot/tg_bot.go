package tg_bot

import (
	CQ_bot "CQ-bot/config"
	"CQ-bot/internial/service/email_service"
	"CQ-bot/internial/service/sheets"
	"CQ-bot/models"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
	"strings"
)

var Teams = make(map[int]models.TeamData)

// TgBot Запуск телеграмм бота
func TgBot(ChanTeam chan models.TeamData) {
	bot, err := tgbotapi.NewBotAPI(CQ_bot.Configs.Tokin_tg)
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	go ExpectationData(bot, ChanTeam)

	//получение апдейтов от тг бота
	for update := range updates {
		if update.CallbackQuery != nil {

			query := fmt.Sprintf("%v", update.CallbackQuery.Data)
			commander := strings.Split(query, "/")
			command := commander[1]

			commandId, err := strconv.Atoi(commander[0])
			if err != nil {
				continue
			}
			teamdata := Teams[commandId]
			if command == "записать" {

				err := email_service.SendMessage(teamdata, true)
				if err != nil {
					continue
				}

				err = sheets.Writedata(teamdata, true)
				if err != nil {
					continue
				}

				msg := fmt.Sprintf("✅Команда *%v* записана на игру *%v*", teamdata.TeamName, teamdata.NameGame)
				editMsg := tgbotapi.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, msg)
				editMsg.ParseMode = tgbotapi.ModeMarkdown
				bot.Send(editMsg)
				delete(Teams, commandId)

			}

			if command == "резерв" {

				err := email_service.SendMessage(teamdata, false)
				if err != nil {
					continue
				}

				err = sheets.Writedata(teamdata, false)
				if err != nil {
					continue
				}

				msg := fmt.Sprintf("⏸Команда *%v* в резерве на игру *%v*", teamdata.TeamName, teamdata.NameGame)
				editMsg := tgbotapi.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, msg)
				editMsg.ParseMode = tgbotapi.ModeMarkdown
				bot.Send(editMsg)
				delete(Teams, commandId)
			}

		}

		//При получении команды "Команды" бот отправляет список команд
		if update.Message != nil && update.Message.Text == "Команды" {
			msg, err := sheets.ParsinTeam()
			if err != nil {
				continue
			}

			msg1 := tgbotapi.NewMessage(update.Message.Chat.ID, msg)
			msg1.ParseMode = tgbotapi.ModeMarkdown
			bot.Send(msg1)
		}

		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

	}

}

// ExpectationData Функция ожидания новых команд
func ExpectationData(bot *tgbotapi.BotAPI, ChanTeam chan models.TeamData) {
	for {
		data, ok := <-ChanTeam
		if !ok {
			continue
		} else {
			if _, ok := Teams[data.ID]; !ok {
				SendMessageTG(bot, data)
				Teams[data.ID] = data
			}
		}
	}
}

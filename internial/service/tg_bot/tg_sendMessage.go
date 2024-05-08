package tg_bot

import (
	CQ_bot "CQ-bot/config"
	"CQ-bot/models"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

// SendMessageTG Функция отправки сообщения о новой команде в тг
func SendMessageTG(bot *tgbotapi.BotAPI, data models.TeamData) {
	bot.Debug = true

	chatID, _ := strconv.ParseInt(CQ_bot.Configs.Chat_ID, 10, 64)

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("✅Записать", fmt.Sprintf("%v/записать", data.ID)),
			tgbotapi.NewInlineKeyboardButtonData("⏸В резерв", fmt.Sprintf("%v/резерв", data.ID)),
		),
	)

	textMsg := fmt.Sprintf("❗Новая запись на игру❗\n\n*команда:* %v\n%v\n%v\n\n%v %v",
		data.TeamName, data.Email, data.Phone, data.DataGame, data.NameGame)

	msg := tgbotapi.NewMessage(chatID, textMsg)
	msg.ReplyMarkup = inlineKeyboard
	msg.ParseMode = tgbotapi.ModeMarkdown

	_, err := bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}
}

package email_service

import (
	CQ_bot "CQ-bot/config"
	"CQ-bot/models"
	"CQ-bot/models/email_msg"
	"fmt"
	gomail "gopkg.in/gomail.v2"
)

func SendMessage(data models.TeamData, flag bool) error {

	var msg_text string

	if flag {
		msg_text = fmt.Sprintf(email_msg.Msg_ok, "%", data.TeamName, data.NameGame, data.DataGame)
	} else {
		msg_text = fmt.Sprintf(email_msg.Msg_res, "%", data.TeamName, data.NameGame, data.DataGame)
	}

	sub := fmt.Sprintf("%v %v", data.NameGame, data.DataGame)

	msg := gomail.NewMessage()
	msg.SetHeader("From", CQ_bot.Configs.Email)
	msg.SetHeader("To", data.Email)
	msg.SetHeader("Subject", sub)

	msg.SetBody("text/html", msg_text)

	mailer := gomail.NewDialer("smtp.yandex.com", 465, CQ_bot.Configs.Email, CQ_bot.Configs.Password)
	err := mailer.DialAndSend(msg)
	if err != nil {
		return err
	}

	return nil
}

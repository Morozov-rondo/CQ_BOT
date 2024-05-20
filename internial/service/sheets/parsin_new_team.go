package sheets

import (
	CQ_bot "CQ-bot/config"
	"CQ-bot/models"
	"fmt"
	"time"
)

// ParsinNewTeam Функция получения новых зарегистрированых команд их отправки в канал
func ParsinNewTeam(chanTeam chan models.TeamData) {
	for {
		teams, err := ParsinSheets(CQ_bot.Configs.CredentialsFile, CQ_bot.Configs.SpreadsheetID, CQ_bot.Configs.RangeName)
		if err != nil {
			continue
		}

		count_id := 2

		for _, team := range teams {

			if team[14] == "false" {

				team_data := models.TeamData{
					ID:       count_id,
					TeamName: fmt.Sprintf("%v", team[0]),
					Email:    fmt.Sprintf("%v", team[1]),
					Phone:    fmt.Sprintf("%v", team[2]),
					DataGame: fmt.Sprintf("%v", team[4]),
					NameGame: fmt.Sprintf("%v", team[13]),
				}

				//Отправка команды в канал
				chanTeam <- team_data
			}
			count_id++
		}

		//Установка времени обращения к Google Sheets
		time.Sleep(15 * time.Second)
	}

}

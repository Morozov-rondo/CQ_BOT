package sheets

import (
	CQ_bot "CQ-bot/config"
	"fmt"
	"strings"
	"time"
)

// ParsinTeam Функция счетчика команд с активными играми
func ParsinTeam() (string, error) {

	counter := make(map[string]int)
	var response string
	teams, err := ParsinSheets(CQ_bot.Configs.CredentialsFile, CQ_bot.Configs.SpreadsheetID, CQ_bot.Configs.RangeName)
	if err != nil {
		return "", err
	}
	for _, team := range teams {

		date_game := fmt.Sprintf("%v", team[4])
		date := strings.Split(date_game, ".")
		date_game = fmt.Sprintf("20%v-%v-%v-19:30:00+07", date[2], date[1], date[0])

		fmt.Println(date_game)

		t, err := time.Parse("2006-01-02-15:04:05-07", date_game)
		if err != nil {
			fmt.Println("тут ошибка")
			continue
		}

		gamestart := t.After(time.Now())
		fmt.Println(gamestart)

		if gamestart && team[14] == "TRUE" {
			name_game := fmt.Sprintf("%v, %v", team[13], team[4])

			if _, ok := counter[name_game]; ok {
				counter[name_game]++
			} else {
				counter[name_game] = 1
			}

		}
	}

	for k, v := range counter {
		result := fmt.Sprintf("📎На *%v*, зарегестрировалось: *%v* команд \n", k, v)
		response += result
	}

	return response, nil
}

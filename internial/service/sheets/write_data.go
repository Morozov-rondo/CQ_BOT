package sheets

import (
	CQ_bot "CQ-bot/config"
	"CQ-bot/models"
	"fmt"
)

func Writedata(data models.TeamData, flag bool) error {
	var values [][]interface{}
	rangeName := fmt.Sprintf("Лист1!O%v:P%v", data.ID, data.ID)
	if flag {
		values = append(values, []interface{}{"TRUE", "Записана"})

		err := WriteDataSheets(CQ_bot.Configs.CredentialsFile, CQ_bot.Configs.SpreadsheetID, rangeName, values)
		if err != nil {
			return err
		}
	} else {
		values = append(values, []interface{}{"TRUE", "В резерве"})

		err := WriteDataSheets(CQ_bot.Configs.CredentialsFile, CQ_bot.Configs.SpreadsheetID, rangeName, values)
		if err != nil {
			return err
		}
	}

	return nil
}

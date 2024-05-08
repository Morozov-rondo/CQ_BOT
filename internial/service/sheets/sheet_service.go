package sheets

import "CQ-bot/internial/GoogleSheetsAPI"

// ParsinSheets Функция получения данных из таблицы
func ParsinSheets(credentialsFile, spreadsheetID, rangeName string) ([][]interface{}, error) {

	client, err := GoogleSheetsAPI.NewGoogleSheetsClient(credentialsFile)
	if err != nil {
		return nil, err
	}

	data, err := client.ReadData(spreadsheetID, rangeName)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func WriteDataSheets(credentialsFile, spreadsheetID, rangeName string, values [][]interface{}) error {

	client, err := GoogleSheetsAPI.NewGoogleSheetsClient(credentialsFile)
	if err != nil {
		return err
	}

	err = client.WriteData(spreadsheetID, rangeName, values)
	if err != nil {
		return err
	}
	return err
}

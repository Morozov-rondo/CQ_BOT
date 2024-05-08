package GoogleSheetsAPI

import (
	"google.golang.org/api/sheets/v4"
)

func (c *GoogleSheetsClient) ReadData(spreadsheetID, rangeName string) ([][]interface{}, error) {

	resp, err := c.service.Spreadsheets.Values.Get(spreadsheetID, rangeName).Do()
	if err != nil {
		return nil, err
	}
	return resp.Values, nil
}

func (c *GoogleSheetsClient) WriteData(spreadsheetID, rangeName string, values [][]interface{}) error {

	_, err := c.service.Spreadsheets.Values.Update(spreadsheetID, rangeName, &sheets.ValueRange{
		Values: values,
	}).ValueInputOption("USER_ENTERED").Do()
	if err != nil {
		return err
	}
	return nil
}

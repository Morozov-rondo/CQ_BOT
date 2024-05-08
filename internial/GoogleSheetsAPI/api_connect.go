package GoogleSheetsAPI

import (
	"context"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type GoogleSheetsClient struct {
	service *sheets.Service
}

// NewGoogleSheetsClient Создание полключения к Google Sheets API
func NewGoogleSheetsClient(credentialsFile string) (*GoogleSheetsClient, error) {
	ctx := context.Background()
	srv, err := sheets.NewService(ctx, option.WithCredentialsFile(credentialsFile))
	if err != nil {
		return nil, err
	}
	return &GoogleSheetsClient{service: srv}, nil
}

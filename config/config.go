//напиши код, который получит данные из .env

package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var Configs Config

type Config struct {
	CredentialsFile string
	SpreadsheetID   string
	RangeName       string
	Tokin_tg        string
	Chat_ID         string
	Email           string
	Password        string
}

func NewConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки файла .env: %v", err)
	}

	Configs.CredentialsFile = os.Getenv("CREDENTAILS_FILE")
	Configs.SpreadsheetID = os.Getenv("SPREAD_SHEET_ID")
	Configs.RangeName = os.Getenv("RANGE_NAME")
	Configs.Tokin_tg = os.Getenv("TOKEN_TG")
	Configs.Chat_ID = os.Getenv("CHAT_ID")
	Configs.Email = os.Getenv("EMAIL")
	Configs.Password = os.Getenv("EMAIL_PASS")

}

package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type SendMailConfig struct {
	Email           string
	Password        string
	Address         string
	Host            string
	DefaultFrom     string
	DefaultFromName string
}

type AppConfig struct {
	MailConfig SendMailConfig
}

func (conf *AppConfig) LoadConfig() {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error: Не удалось загрузить файл конфигурации: %v", err)
		os.Setenv("config_init", "0")
	} else {
		os.Setenv("config_init", "1")
	}

	conf.MailConfig.Email = os.Getenv("SENDMAIL_EMAIL")
	conf.MailConfig.Password = os.Getenv("SENDMAIL_PASSWORD")
	conf.MailConfig.Address = os.Getenv("SENDMAIL_ADDRESS")
	conf.MailConfig.Host = os.Getenv("SENDMAIL_HOST")
	conf.MailConfig.DefaultFrom = os.Getenv("SENDMAIL_DEFAULT_FROM")
	conf.MailConfig.DefaultFromName = os.Getenv("SENDMAIL_DEFAULT_FROM_NAME")
}

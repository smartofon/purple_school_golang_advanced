package verify

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/smtp"
	"purple_school_golang_advanced/3_validation_api/configs"
	"purple_school_golang_advanced/3_validation_api/internal/storage"

	"github.com/jordan-wright/email"
)

func SendVerify(login string, server configs.SendMailConfig) {

	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", server.DefaultFromName, server.DefaultFrom)
	e.To = []string{login}
	e.Subject = "Восстановление пароля"

	h := createHash(10)

	url := fmt.Sprintf("http://localhost:8081/verify/%s", h)

	e.Text = []byte(fmt.Sprintf("Используйте для восстановления пароля перейдите по ссылке: %s", url))
	e.HTML = []byte(fmt.Sprintf("<b>%s</b>", fmt.Sprintf("Используйте для восстановления перейдите по ссылке: <a href=\"%s\">%s</a>", url, url)))

	err := e.Send(server.Address, smtp.PlainAuth("", server.Email, server.Password, server.Host))
	if err != nil {
		log.Fatalf("Error: Ошибка при отправке сообщения: %v\n %+v", err, e)
	}

	if storage.GlobalStorage != nil {
		fmt.Println("Сохраняем хранилище")
		storage.GlobalStorage.Set(h, login)
		storage.GlobalStorage.Save()
	}
}

func VerifyHash(hash string) bool {
	_, ok := storage.GlobalStorage.Get(hash)
	return ok
}

func DeleteHash(hash string) {
	storage.GlobalStorage.Delete(hash)
	storage.GlobalStorage.Save()
}

func createHash(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

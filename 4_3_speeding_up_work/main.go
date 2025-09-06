package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	n := 10

	fmt.Printf("Выполняем %d соединений\n", n)

	fmt.Println("Без горутин")

	t := time.Now()

	for i := 0; i < n; i++ {
		getHttpCode(i)
	}

	fmt.Printf("Время исполнения без горутин: %v\n", time.Since(t))

	fmt.Println("С горутинами")

	t = time.Now()

	for i := 0; i < n; i++ {
		go getHttpCode(i)
	}

	// горутины могут отработать быстрее,
	// но пока ограничим исполнения в разы по отношению к предыдущему участку
	// с использованием задержки исполнения
	time.Sleep(time.Second)

	fmt.Printf("Время исполнения sleep: %v\n", time.Since(t))

}

func getHttpCode(n int) {
	resp, err := http.Get("https://google.com")
	if err == nil {
		fmt.Printf("Код ответа %d: %d\n", n, resp.StatusCode)
	} else {
		fmt.Printf("Ошибка: %s", err.Error())
	}
}

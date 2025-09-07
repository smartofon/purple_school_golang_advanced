package main

import (
	"fmt"
	"os"
	"purple_school_golang_advanced/6_lesson/configs"
	"strings"
)

func main() {
	config := configs.LoadConfig()
	fmt.Printf("%+v", config)

	for _, env := range os.Environ() {
		s := strings.Split(env, "=")
		fmt.Println(s[0], s[1])
	}
}

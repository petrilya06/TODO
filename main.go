package main

import (
	"log"
	"tgbot_todolist/bot"
)

func main() {
	b, err := bot.NewBot()
	if err != nil {
		log.Fatal(err)
	}

	b.Run()
}

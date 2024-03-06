package main

import (
	"aslon1213/customer_support_bot/pkg/app"
)

func main() {
	app := app.New()
	go app.LoadUsagesFromRedisToDatabase()
	app.Run()
	defer app.Close()
}

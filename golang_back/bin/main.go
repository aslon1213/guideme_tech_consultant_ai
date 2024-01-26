package main

import "aslon1213/customer_support_bot/internal/app"

func main() {
	app := app.New()
	app.Run()
	defer app.Close()
}

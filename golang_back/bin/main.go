package main

import "aslon1213/customer_support_bot/pkg/app"

func main() {
	app := app.New()
	app.Run()
	defer app.Close()
}

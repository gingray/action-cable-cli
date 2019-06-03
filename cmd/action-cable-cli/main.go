package main

import (
	"action-cable-cli/client"
	"action-cable-cli/ui"
)

func main() {
	app := ui.BuildUI(client.GetInstance())

	if err := app.Run(); err != nil {
		panic(err)
	}

}

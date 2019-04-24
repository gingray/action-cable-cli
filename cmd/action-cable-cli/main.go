package main

import (
	"action-cable-cli/client"
	"action-cable-cli/ui"
)

func main() {
	config := client.Config{}
	// configWs := &client.Config{Url: "wss://echo.websocket.org"}
	// client, err := client.NewClient(configWs)
	// client.WriteMessage("tatatata")
	// if err != nil {
	// 	return
	// }
	// go client.ReadLoop()
	// go func() {
	// 	for {
	// 		buffer := <-client.ChReadBuffer
	// 		app.QueueUpdateDraw(func() {
	// 			fmt.Fprintf(config.StatusText, "%s", buffer)
	// 		})

	// 	}
	// }()

	app := ui.BuildUI(&config)

	if err := app.Run(); err != nil {
		panic(err)
	}

}

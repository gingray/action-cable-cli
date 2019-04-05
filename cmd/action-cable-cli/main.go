package main

import (
	"action-cable-cli/client"
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	inputField := tview.NewInputField().
		SetLabel("Enter a number: ").
		SetFieldWidth(10).
		SetAcceptanceFunc(tview.InputFieldInteger).
		SetDoneFunc(func(key tcell.Key) {
			fmt.Println(key)
			app.Stop()
		})
	if err := app.SetRoot(inputField, true).SetFocus(inputField).Run(); err != nil {
		panic(err)
	}

	config := &client.Config{Url: "wss://echo.websocket.org"}
	client, err := client.NewClient(config)
	if err != nil {
		return
	}
	fmt.Printf("%v", &client)
}

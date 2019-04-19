package main

import (
	"action-cable-cli/client"
	"fmt"
	"net/url"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	config := client.Config{}
	config.UrlField = tview.NewInputField().
		SetLabel("WS URL: ").
		SetFieldWidth(100).
		SetDoneFunc(func(key tcell.Key) {
			_, err := url.ParseRequestURI(config.UrlField.GetText())
			if err != nil {
				config.UrlField.SetLabel("WS URL(error): ")
				config.UrlField.SetLabelColor(tcell.ColorOrangeRed)
				return
			}
			config.UrlField.SetLabelColor(tcell.ColorPaleGreen)
			config.UrlField.SetLabel("WS URL: ")

		})
	config.StatusText = tview.NewTextView()
	config.StatusText.SetChangedFunc(func() {
		app.Draw()
	})
	config.StatusText.SetWrap(true)

	grid := tview.NewGrid().
		SetRows(-1, -1, -1).
		SetColumns(-1, -1, -1).
		AddItem(config.UrlField, 0, 0, 1, 2, 0, 0, true)
	grid.SetBackgroundColor(tview.Styles.PrimitiveBackgroundColor)

	grid.AddItem(config.StatusText, 0, 2, 1, 1, 0, 0, false)

	configWs := &client.Config{Url: "wss://echo.websocket.org"}
	client, err := client.NewClient(configWs)
	client.WriteMessage("tatatata")
	if err != nil {
		return
	}
	go client.ReadLoop()
	go func() {
		for {
			buffer := <-client.ChReadBuffer
			app.QueueUpdateDraw(func() {
				fmt.Fprintf(config.StatusText, "%s", buffer)
			})

		}
	}()

	if err := app.SetRoot(grid, true).SetFocus(grid).Run(); err != nil {
		panic(err)
	}

}

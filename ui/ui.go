package ui

import (
	"action-cable-cli/client"
	"action-cable-cli/helpers"
	"net/url"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)


type UI struct {
	Field      *tview.InputField
	Input      *tview.InputField
	SendBtn    *tview.Button
	ConnectBtn *tview.Button
}

var mainUI *UI

func BuildUI(config *client.Config) *tview.Application {
	elements := []tview.Primitive{}
	currentFocus := 0
	app := tview.NewApplication()
	grid := tview.NewGrid().
		SetRows(3, 3, 1, -1).
		SetColumns(-1, -1, -1).
		SetGap(0, 1)
	mainUI = &UI{}
	grid.SetBackgroundColor(tview.Styles.PrimitiveBackgroundColor)
	elements = append(elements, createField(grid))
	elements = append(elements, createMethodInput(grid))
	elements = append(elements, createSendBtn(grid))
	elements = append(elements, createConnectBtn(grid))

	// elements = append(elements, createOutLogField(grid, app))
	app.SetRoot(grid, true).SetFocus(grid)
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyTab {
			currentFocus += 1
			if currentFocus >= len(elements) {
				currentFocus = 0
			}
			app.SetFocus(elements[currentFocus])
			//set focus different UI items
			//panic("")
		}
		return event
	})
	cl := client.GetInstance()
	go mainUI.UpdateUILoop(cl.UIChan)
	return app
}

func (self *UI ) UpdateUILoop(ch chan helpers.UIMsg) {
for uiMsg := range ch {
	if self.Field != nil && uiMsg.MsgType== helpers.UI_INFO {
		self.Field.SetText(uiMsg.Msg)
	}
}
}

func createField(root *tview.Grid) tview.Primitive {
	inputField := tview.NewInputField().
		SetLabel("WS URL: ").
		SetFieldWidth(100)
	inputField.SetDoneFunc(func(key tcell.Key) {
		_, err := url.ParseRequestURI(inputField.GetText())
		if err != nil {
			inputField.SetLabel("WS URL(error): ")
			inputField.SetLabelColor(tcell.ColorOrangeRed)
			return
		}
		inputField.SetLabelColor(tcell.ColorPaleGreen)
		inputField.SetLabel("WS URL: ")
	})
	inputField.SetLabelColor(tcell.ColorPaleGreen)
	inputField.SetLabel("WS URL: ")
	root.AddItem(inputField, 0, 0, 1, 2, 0, 0, true)
	return inputField
}

func createMethodInput(root *tview.Grid) tview.Primitive {
	methodInput := tview.NewInputField().
		SetLabel("Method to call: ").
		SetFieldWidth(100)
	root.AddItem(methodInput, 1, 0, 1, 2, 0, 0, true)
	return methodInput
}

func createSendBtn(root *tview.Grid) tview.Primitive {
	sendBtn := tview.NewButton("Send")
	sendBtn.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		return nil
	})

	root.AddItem(sendBtn, 2, 0, 1, 1, 0, 0, true)
	return sendBtn
}

func createConnectBtn(root *tview.Grid) tview.Primitive {
	connectBtn := tview.NewButton("Connect")
	connectBtn.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter {
			cl := client.GetInstance()
			cl.Connect()
		}
		return nil
	})

	root.AddItem(connectBtn, 2, 1, 1, 1, 0, 0, true)
	return connectBtn
}

func createOutLogField(root *tview.Grid, app *tview.Application) tview.Primitive {
	statusText := tview.NewTextView()
	statusText.SetChangedFunc(func() {
		app.Draw()
	})
	statusText.SetWrap(true)
	root.AddItem(statusText, 3, 0, 1, 3, 0, 0, false)
	return statusText
}

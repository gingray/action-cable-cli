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
	InputField      *tview.InputField
	SendBtn    *tview.Button
	ConnectBtn *tview.Button
}

var mainUI *UI

func BuildUI(cl *client.Client) *tview.Application {
	elements := []tview.Primitive{}
	currentFocus := 0
	app := tview.NewApplication()
	grid := tview.NewGrid().
		SetRows(3, 3, 1, -1).
		SetColumns(-1, -1, -1).
		SetGap(0, 1)
	mainUI = &UI{}
	grid.SetBackgroundColor(tview.Styles.PrimitiveBackgroundColor)
	elements = append(elements, mainUI.createField(grid))
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
	go mainUI.UpdateUILoop(cl.UIChan)
	return app
}

func (self *UI ) UpdateUILoop(ch chan helpers.UIMsg) {
for uiMsg := range ch {
	if self.InputField != nil && uiMsg.MsgType== helpers.UI_INFO {
		self.InputField.SetText(uiMsg.Msg)
	}
}
}

func (self *UI) createField(root *tview.Grid) tview.Primitive {
	self.InputField = tview.NewInputField().
		SetLabel("WS URL: ").
		SetFieldWidth(100)
	self.InputField.SetDoneFunc(func(key tcell.Key) {
		_, err := url.ParseRequestURI(self.InputField.GetText())
		if err != nil {
			self.InputField.SetLabel("WS URL(error): ")
			self.InputField.SetLabelColor(tcell.ColorOrangeRed)
			return
		}
		self.InputField.SetLabelColor(tcell.ColorPaleGreen)
		self.InputField.SetLabel("WS URL: ")
	})
	self.InputField.SetLabelColor(tcell.ColorPaleGreen)
	self.InputField.SetLabel("WS URL: ")
	root.AddItem(self.InputField, 0, 0, 1, 2, 0, 0, true)
	return self.InputField
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

package ui

import (
	"action-cable-cli/client"
	"action-cable-cli/helpers"
	"net/url"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)


type UI struct {
	MethodField *tview.InputField
	UrlField    *tview.InputField
	SendBtn     *tview.Button
	ConnectBtn  *tview.Button
	StatusText *tview.TextView
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
	elements = append(elements, mainUI.createUrlInput(grid))
	elements = append(elements, mainUI.createMethodInput(grid))
	elements = append(elements, mainUI.createSendBtn(grid))
	elements = append(elements, mainUI.createConnectBtn(grid))
	mainUI.createOutLogField(grid, app)

	// elements = append(elements, createOutLogField(grid, app))
	app.SetRoot(grid, true).SetFocus(grid)
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyTab {
			currentFocus += 1
			if currentFocus >= len(elements) {
				currentFocus = 0
			}
			app.SetFocus(elements[currentFocus])
		}
		return event
	})
	go mainUI.UpdateUILoop(cl.UIChan)
	return app
}

func (self *UI ) UpdateUILoop(ch chan helpers.UIMsg) {
for uiMsg := range ch {
	if self.StatusText != nil && uiMsg.MsgType== helpers.UI_INFO {
		self.StatusText.SetText(uiMsg.Msg)
	}
}
}

func (self *UI) createUrlInput(root *tview.Grid) tview.Primitive {
	self.UrlField = tview.NewInputField().
		SetLabel("WS URL: ").
		SetFieldWidth(100)
	self.UrlField.SetDoneFunc(func(key tcell.Key) {
		_, err := url.ParseRequestURI(self.UrlField.GetText())
		if err != nil {
			self.UrlField.SetLabel("WS URL(error): ")
			self.UrlField.SetLabelColor(tcell.ColorOrangeRed)
			return
		}
		self.UrlField.SetLabelColor(tcell.ColorPaleGreen)
		self.UrlField.SetLabel("WS URL: ")
		cl :=client.GetInstance()
		cl.Config.Url = self.UrlField.GetText()
	})
	self.UrlField.SetLabelColor(tcell.ColorPaleGreen)
	self.UrlField.SetLabel("WS URL: ")
	root.AddItem(self.UrlField, 0, 0, 1, 2, 0, 0, true)
	return self.UrlField
}

func (self * UI) createMethodInput(root *tview.Grid) tview.Primitive {
	self.MethodField = tview.NewInputField().
		SetLabel("Method to call: ").
		SetFieldWidth(100)
	root.AddItem(self.MethodField, 1, 0, 1, 2, 0, 0, true)
	return self.MethodField
}

func (self *UI) createSendBtn(root *tview.Grid) tview.Primitive {
	self.SendBtn = tview.NewButton("Send")
	self.SendBtn.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		return nil
	})

	root.AddItem(self.SendBtn, 2, 0, 1, 1, 0, 0, true)
	return self.SendBtn
}

func (self *UI) createConnectBtn(root *tview.Grid) tview.Primitive {
	self.ConnectBtn = tview.NewButton("Connect")
	self.ConnectBtn.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter {
			cl := client.GetInstance()
			cl.Connect()
		}
		return nil
	})

	root.AddItem(self.ConnectBtn, 2, 1, 1, 1, 0, 0, true)
	return self.ConnectBtn
}

func (self *UI) createOutLogField(root *tview.Grid, app *tview.Application) tview.Primitive {
	self.StatusText = tview.NewTextView()
	self.StatusText.SetChangedFunc(func() {
		app.Draw()
	})
	self.StatusText.SetWrap(true)
	root.AddItem(self.StatusText, 3, 0, 1, 3, 0, 0, false)
	return self.StatusText
}

package client

import "github.com/rivo/tview"

type Config struct {
	Url        string
	UrlField   *tview.InputField
	StatusText *tview.TextView
}

package client

import "os"

type Config struct {
	Url     string
	Method  string
	Payload string
}
//https://bestvpn.org/html5demos/web-socket/

func NewConfig() *Config {
	url:= os.Getenv("ACTION_CABLE_URL")
	return &Config{Url:url}
}

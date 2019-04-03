package client

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	config *Config
	conn   *websocket.Conn
}

func NewClient(config *Config) (client *Client, err error) {
	conn, _, err := websocket.DefaultDialer.Dial(config.Url, nil)
	if err != nil {
		return nil, err
	}
	return &Client{conn: conn, config: config}, nil
}

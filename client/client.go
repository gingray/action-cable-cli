package client

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	config *Config
	conn   *websocket.Conn
}

func NewClient(config *Config) (client *Client, err error) {
	conn, resp, err := websocket.DefaultDialer.Dial(config.Url, nil)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
		fmt.Println("Connection success")
	}
	if err != nil {
		return nil, err
	}
	return &Client{conn: conn, config: config}, nil
}

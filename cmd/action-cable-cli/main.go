package main

import (
	"action-cable-cli/client"
	"fmt"
)

func main() {
	config := &client.Config{Url: "wss://echo.websocket.org"}
	client, _ := client.NewClient(config)
	fmt.Printf("%v", &client)
}

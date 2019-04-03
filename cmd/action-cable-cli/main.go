package main

import (
	"action-cable-cli/client"
	"fmt"
)

func main() {
	client := client.Client{}
	fmt.Printf("%v", &client)
}

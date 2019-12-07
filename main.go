package main

import (
	"GoTicket/general"
	"fmt"
)

func main() {
	config := general.TicketConfig()
	fmt.Println(config)
}


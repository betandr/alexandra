package main

import (
	"github.com/betandr/alexandra/client"
	"github.com/betandr/alexandra/server"
)

func main() {
	go server.Start()
	client.Start()
}

package main

import (
	server "github.com/PedroSMarcal/Naval_Battle/Server"
)

func initialize() {
	go server.Start(":8080")

	go server.Start(":8081")

}

func main() {

	defer initialize()
}

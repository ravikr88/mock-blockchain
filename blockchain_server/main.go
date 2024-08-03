package main

import (
	"flag"
	"fmt"
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	port := flag.Uint("port", 8000, "TCP Port Number for Blockchain Server")
	flag.Parse()
	fmt.Println(*port)
	app := NewBlockchainServer(uint16(*port))
	app.Run()
}

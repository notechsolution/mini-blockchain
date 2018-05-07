package main

import (
	"log"
	"github.com/joho/godotenv"
	"./chain"
	"./server"
)


func main() {
	err := godotenv.Load("./networkVersion/config.env")
	if err != nil {
		log.Fatal(err)
	}
	go networkChain.InitialGenesisBlock();
	tcpServer.Run();
}




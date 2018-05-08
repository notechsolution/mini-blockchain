package main

import (
	"log"
	"github.com/joho/godotenv"
	"./chain"
	"./server"
)


func main() {
	err := godotenv.Load("./httpWithMinerVersion/config.env")
	if err != nil {
		log.Fatal(err)
	}
	go minerChain.InitialGenesisBlock();
	log.Fatal(server.Run());
}




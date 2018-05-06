package main

import (
	"log"
	"github.com/joho/godotenv"
	"./chain"
	"./server"
)


func main() {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal(err)
	}
	go chain.InitialGenesisBlock();
	log.Fatal(server.Run());
}




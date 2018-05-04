package main

import (
	"./model"
	"os"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"time"
	"encoding/json"
	"io"
)

type Block model.Block;

var Blockchain []Block;

func main() {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal(err)
	}
	go initialGenesisBlock();
	run();
}

func initialGenesisBlock() {
	var genesis Block;
	genesis.Index = 0;

}

func replaceChain(newBlocks [] Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks;
	}
}

func run() error {
	httpAddr := os.Getenv("ADDR");
	log.Println("Listening on ", httpAddr);
	mux := makeMuxRouter();
	s := &http.Server{
		Addr:":" + httpAddr,
		Handler: mux,
		ReadTimeout: 10	* time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1<<20,
	}

	if err := s.ListenAndServe(); err != nil {
		return err;
	}

	return nil;
}

func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter();
	muxRouter.HandleFunc("/", handleGetBlockchain).Methods("GET");
	return muxRouter;
}
func handleGetBlockchain(writer http.ResponseWriter, request *http.Request) {
	log.Printf("[%s]Receiced request '%s' ", "handleGetBlockchain", request.RequestURI);
	bytes, err := json.MarshalIndent(Blockchain,"", " ");
	if err != nil {
		http.Error(writer,err.Error(), http.StatusInternalServerError);
		return;
	}
	io.WriteString(writer, string(bytes));
}

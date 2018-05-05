package main

import (
	"./model"
	"./generator"
	"os"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"time"
	"encoding/json"
	"io"
	"github.com/davecgh/go-spew/spew"
)


var Blockchain []model.Block;

func main() {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal(err)
	}
	go initialGenesisBlock();
	log.Fatal(run());
}

func initialGenesisBlock() {
	var genesis model.Block;
	genesis.Index = 0;
	genesis.BPM = 0;
	genesis.Timestamp = time.Now().String();
	genesis.PrevHash = os.Getenv("PARENTHASH");
	genesis.Hash = generator.CalculateHash(genesis);
	Blockchain = append(Blockchain, genesis);
	spew.Dump(genesis)
}

func replaceChain(newBlocks [] model.Block) {
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
	muxRouter.HandleFunc("/getAllBlocks", handleGetBlockchain).Methods("GET");
	muxRouter.HandleFunc("/create", handleCreateBlockchain).Methods("POST");
	return muxRouter;
}

type Message struct {
	BPM int;
}

func handleCreateBlockchain(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body);
	var m Message;
	if err:=decoder.Decode(&m); err !=nil {
		responseWithJSON(writer, request, http.StatusBadRequest, request.Body);
		return
	}
	defer request.Body.Close();
	previousBlock := Blockchain[len(Blockchain)-1];
	newBlock, err:=generator.GenerateBlock(previousBlock, m.BPM);
	if err !=nil {
		responseWithJSON(writer, request, http.StatusInternalServerError, m);
	}

	if generator.IsValidBlock(newBlock,previousBlock) {
		newBlockchain := append(Blockchain, newBlock);
		replaceChain(newBlockchain);
		spew.Dump(newBlock);
	}

	responseWithJSON(writer, request, http.StatusCreated, newBlock);
}
func responseWithJSON(writer http.ResponseWriter, request *http.Request, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", " ");
	if err!=nil {
		writer.WriteHeader(http.StatusInternalServerError);
		writer.Write([]byte("HTTP 500: Internal Server Error\n" +err.Error()));
		return;
	}
	writer.WriteHeader(code);
	writer.Write(response);
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

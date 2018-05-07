package server

import (
	"net/http"
	"../chain"
	"encoding/json"
	"github.com/gorilla/mux"
	"os"
	"log"
	"time"
	"io"
)

func Run() error {
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
	newBlock, err := chain.CreateNewBlock(m.BPM);
	if err !=nil {
		responseWithJSON(writer, request, http.StatusInternalServerError, m);
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
	bytes, err := json.MarshalIndent(chain.GetAllBlocks(),"", " ");
	if err != nil {
		http.Error(writer,err.Error(), http.StatusInternalServerError);
		return;
	}
	io.WriteString(writer, string(bytes));
}

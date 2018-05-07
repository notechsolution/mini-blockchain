package tcpServer

import (
	"os"
	"net"
	"log"
	"io"
	"bufio"
	"strconv"
	"../chain"
	"../../model"
	"github.com/davecgh/go-spew/spew"
	"time"
)

var bcServer chan []model.Block;

func Run() {
	bcServer = make(chan [] model.Block);
	tcpPort := os.Getenv("TCP_PORT");
	server, err := net.Listen("tcp", ":"+tcpPort);
	if err != nil {
		log.Fatal(err);
	}

	for {
		conn, err := server.Accept();
		if err != nil {
			log.Fatal(err);
		}
		go handleConnection(conn);
	}
}
func handleConnection(conn net.Conn) {
	defer conn.Close();
	io.WriteString(conn, "Please enter BPM:");
	scanner := bufio.NewScanner(conn);
	go func() {
		for scanner.Scan() {
			bpm, err := strconv.Atoi(scanner.Text());
			if err != nil {
				log.Println(err);
				continue;
			}

			newBlock, err := networkChain.CreateNewBlock(bpm);
			if err != nil {
				log.Println(err);
				continue;
			}
			spew.Dump(newBlock);
			log.Println("ready to save into channel");
			bcServer <- networkChain.GetAllBlocks();
			io.WriteString(conn, "\nPlease enter BPM:");
		}
	}();
	go broadcast(conn);

	for _ = range bcServer {
		spew.Dump(networkChain.GetAllBlocks())
	}
}
func broadcast(conn net.Conn) {
	for {
		time.Sleep(10 * time.Second);
		output := networkChain.GetSyncOutput();
		io.WriteString(conn, "received from broadcast,"+output);
	}
}

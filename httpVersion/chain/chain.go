package chain

import (
	"time"
	"os"
	"github.com/davecgh/go-spew/spew"
	"../../generator"
	"../../model"
)

var blockchain []model.Block;

func InitialGenesisBlock() {
	var genesis model.Block;
	genesis.Index = 0;
	genesis.BPM = 0;
	genesis.Timestamp = time.Now().String();
	genesis.PrevHash = os.Getenv("PARENTHASH");
	genesis.Hash = generator.CalculateHash(genesis);
	blockchain = append(blockchain, genesis);
	spew.Dump(genesis)
}


func ReplaceChain(newBlocks [] model.Block) {
	if len(newBlocks) > len(blockchain) {
		blockchain = newBlocks;
	}
}

func CreateNewBlock(BPM int) (model.Block, error) {
	previousBlock := blockchain[len(blockchain)-1];
	newBlock, err:=generator.GenerateBlock(previousBlock, BPM);
	if err !=nil {
		return newBlock, err;
	}

	if generator.IsValidBlock(newBlock,previousBlock) {
		newBlockchain := append(blockchain, newBlock);
		ReplaceChain(newBlockchain);
		spew.Dump(newBlock);
	}
	return newBlock, err;
}

func GetAllBlocks()[]model.Block{
	return blockchain;
}

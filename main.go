package main

import (
	"./model"
)

type Block model.Block;

var Blockchain []Block;

func main() {

}

func replaceChain(newBlocks [] Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks;
	}
}

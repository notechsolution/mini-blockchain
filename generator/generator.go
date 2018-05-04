package generator

import (
	"../model"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
);

type Block model.Block;

func CalculateHash(block Block) string {
	content := strconv.Itoa(block.BPM) + strconv.Itoa(block.Index) + block.Timestamp + block.PrevHash;
	sha := sha256.New();
	sha.Write([]byte(content));
	hashed := sha.Sum(nil);
	return hex.EncodeToString(hashed);
}

func GenerateBlock(oldBlock Block, BPM int) (Block, error) {
	var block Block;
	block.BPM = BPM;
	block.Index = oldBlock.Index + 1;
	block.PrevHash = oldBlock.Hash;
	block.Timestamp = time.Now().String();
	block.Hash = CalculateHash(block);
	return block, nil;
}

func isValidBlock(newBlock, oldBlock Block) bool {
	if newBlock.Index != oldBlock.Index+1 {
		return false;
	}

	if newBlock.PrevHash != oldBlock.Hash {
		return false;
	}

	if CalculateHash(newBlock) != newBlock.Hash {
		return false;
	}
	return true;
}

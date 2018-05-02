package generator

import (
	"../model"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
);

func CalculateHash(block model.Block) string {
	content := strconv.Itoa(block.BPM) + strconv.Itoa(block.Index) + block.Timestamp + block.PrevHash;
	fmt.Printf("content is %s, %d, %d\n", content,block.BPM, block.Index);
	sha := sha256.New();
	sha.Write([]byte(content));
	hashed := sha.Sum(nil);
	return hex.EncodeToString(hashed);
}

func GenerateBlock(oldBlock model.Block, BPM int) (model.Block, error) {
	var block model.Block;
	block.BPM = BPM;
	block.Index = oldBlock.Index +1;
	block.PrevHash = oldBlock.Hash;
	block.Timestamp = time.Now().String();
	block.Hash = CalculateHash(block);
	return block, nil;
}

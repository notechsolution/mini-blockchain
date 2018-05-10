package generator

import (
	"../model"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
	"fmt"
	"strings"
	"log"
);

func CalculateHash(block model.Block) string {
	content := strconv.Itoa(block.BPM) + strconv.Itoa(block.Index) + block.Timestamp + block.PrevHash + block.Nonce;
	sha := sha256.New();
	sha.Write([]byte(content));
	hashed := sha.Sum(nil);
	return hex.EncodeToString(hashed);
}

func GenerateBlock(oldBlock model.Block, BPM int) (model.Block, error) {
	var block model.Block;
	block.BPM = BPM;
	block.Index = oldBlock.Index + 1;
	block.PrevHash = oldBlock.Hash;
	block.Timestamp = time.Now().String();
	block.Hash = CalculateHash(block);
	return block, nil;
}

func GenerateBlockWithDifficulty(oldBlock model.Block, BPM int, difficulty int) (model.Block, error) {
	var block model.Block;
	block.BPM = BPM;
	block.Index = oldBlock.Index + 1;
	block.PrevHash = oldBlock.Hash;
	block.Timestamp = time.Now().String();
	mining(&block, difficulty)

	block.Hash = CalculateHash(block);
	return block, nil;
}

func mining(block *model.Block, difficulty int) {
	for i := 0; ; i++ {
		block.Nonce = fmt.Sprintf("%x", i);
		hash := CalculateHash(*block);
		if !isValidHash(hash, difficulty) {
			log.Printf("%s	do more work\n", hash);
			time.Sleep(50 * time.Millisecond);
			continue;
		} else {
			log.Printf("%s	well done", hash);
			break;
		}
	}
}

func isValidHash(hash string, difficulty int) bool {
	prefix := strings.Repeat("0", difficulty);
	return strings.HasPrefix(hash, prefix);
}

func IsValidBlock(newBlock, oldBlock model.Block) bool {
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

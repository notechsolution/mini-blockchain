package generator

import (
	"../model"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
);

func CalculateHash(block model.Block) string {
	content := strconv.Itoa(block.BPM) + strconv.Itoa(block.Index) + block.Timestamp + block.PrevHash;
	fmt.Printf("content is %s, %d, %d\n", content,block.BPM, block.Index);
	sha := sha256.New();
	sha.Write([]byte(content));
	hashed := sha.Sum(nil);
	return hex.EncodeToString(hashed);
}

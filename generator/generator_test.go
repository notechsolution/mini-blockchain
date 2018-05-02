package generator

import (
	"testing"

	"../model"
)

func TestCalculateHash(t *testing.T) {
	block :=  model.Block{
		BPM:       10,
		Index:     6,
		Timestamp: "2334455666",
		PrevHash:  "omygodiamaprevhash",
	} // 1062334455666omygodiamaprevhash

	want := "07948e6daa751dd9295e181d58311c15f67fca504f8ad6e0600814666cc6d9b7";
	if got := CalculateHash(block); got != want {
		t.Errorf("CalculateHash() = %v, want %v", got, want)
	}
}

package generator

import (
	"reflect"
	"testing"

	"../model"
	"go/doc"
)

func TestCalculateHash(t *testing.T) {
	type args struct {
		block model.Block
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "generate hash string correctly",
			args: args{
				block: model.Block{
					BPM:       10,
					Index:     6,
					Timestamp: "2334455666",
					PrevHash:  "omygodiamaprevhash",
				},
			},
			want: "07948e6daa751dd9295e181d58311c15f67fca504f8ad6e0600814666cc6d9b7",
		}, // can add more test here
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateHash(tt.args.block); got != tt.want {
				t.Errorf("CalculateHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

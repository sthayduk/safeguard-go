package main

import (
	"fmt"

	safeguard "github.com/sthayduk/safeguard-go"
	"github.com/sthayduk/safeguard-go/examples/common"
)

func main() {
	err := common.InitClient()
	if err != nil {
		panic(err)
	}

	partitions, err := safeguard.GetAssetPartitions(safeguard.Filter{})
	if err != nil {
		panic(err)
	}

	for _, partition := range partitions {
		jsonStr, err := partition.ToJson()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", jsonStr)
	}
}

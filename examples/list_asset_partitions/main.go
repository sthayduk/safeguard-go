package main

import (
	"fmt"

	"github.com/sthayduk/safeguard-go"
	"github.com/sthayduk/safeguard-go/examples/common"
)

func main() {
	sgc, err := common.InitClient()
	if err != nil {
		panic(err)
	}

	partitions, err := sgc.GetAssetPartitions(safeguard.Filter{})
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

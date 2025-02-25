package main

import (
	"fmt"

	"github.com/sthayduk/safeguard-go/client"
	"github.com/sthayduk/safeguard-go/examples/common"
	"github.com/sthayduk/safeguard-go/models"
)

func main() {
	sgc, err := common.InitClient()
	if err != nil {
		panic(err)
	}

	partitions, err := models.GetAssetPartitions(sgc, client.Filter{})
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

package main

import (
	"context"
	"fmt"
	"time"

	safeguard "github.com/sthayduk/safeguard-go"
	"github.com/sthayduk/safeguard-go/client"
	"github.com/sthayduk/safeguard-go/examples/common"
)

func main() {
	err := common.InitClient()
	if err != nil {
		panic(err)
	}

	assetAccount, err := safeguard.GetAssetAccount(18, client.Fields{})
	if err != nil {
		panic(err)
	}

	checkPasswordTask, err := assetAccount.CheckPassword()
	if err != nil {
		panic(err)
	}

	// Check the task state
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()
	state, err := checkPasswordTask.CheckTaskState(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Password Check State: ", state)
}

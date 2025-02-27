package main

import (
	"context"
	"fmt"
	"time"

	. "github.com/sthayduk/safeguard-go"
	"github.com/sthayduk/safeguard-go/client"
	"github.com/sthayduk/safeguard-go/examples/common"
)

func main() {
	sgc, err := common.InitClient()
	if err != nil {
		panic(err)
	}

	assetAccount, err := GetAssetAccount(sgc, 249, client.Fields{})
	if err != nil {
		panic(err)
	}

	changePasswordTask, err := assetAccount.ChangePassword()
	if err != nil {
		panic(err)
	}

	// Wait for the task to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	state, err := changePasswordTask.CheckTaskState(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Password Change State: ", state)
}

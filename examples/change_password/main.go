package main

import (
	"fmt"

	"github.com/sthayduk/safeguard-go/examples/common"
	"github.com/sthayduk/safeguard-go/src/client"
	"github.com/sthayduk/safeguard-go/src/models"
)

func main() {
	sgc, err := common.InitClient()
	if err != nil {
		panic(err)
	}

	assetAccount, err := models.GetAssetAccount(sgc, 18, client.Fields{})
	if err != nil {
		panic(err)
	}

	changePasswordTask, err := assetAccount.ChangePassword(sgc)
	if err != nil {
		panic(err)
	}

	state, err := changePasswordTask.CheckTaskState(sgc)
	if err != nil {
		panic(err)
	}
	fmt.Println("Password Change State: ", state)
}

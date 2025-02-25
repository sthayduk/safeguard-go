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

	assetAccount, err := models.GetAssetAccount(sgc, 18, client.Fields{})
	if err != nil {
		panic(err)
	}

	checkPasswordTask, err := assetAccount.CheckPassword()
	if err != nil {
		panic(err)
	}

	state, err := checkPasswordTask.CheckTaskState()
	if err != nil {
		panic(err)
	}
	fmt.Println("Password Check State: ", state)
}

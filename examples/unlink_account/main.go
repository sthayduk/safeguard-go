package main

import (
	"fmt"

	"github.com/sthayduk/safeguard-go/examples/common"
	"github.com/sthayduk/safeguard-go/models"
)

func main() {
	sgc, err := common.InitClient()
	if err != nil {
		panic(err)
	}

	// Id 76 = "Stefan Hayduk"
	user, err := models.GetUser(sgc, 76, nil)
	if err != nil {
		panic(err)
	}

	// Id 133 = "da-andresen"
	policyAccount, err := models.GetPolicyAccount(sgc, 133, nil)
	if err != nil {
		panic(err)
	}

	linkedAccount, err := models.RemoveLinkedAccounts(sgc, user, []models.PolicyAccount{policyAccount})
	if err != nil {
		panic(err)
	}

	// Id 134 = "sa-andresen"
	policyAccount, err = models.GetPolicyAccount(sgc, 134, nil)
	if err != nil {
		panic(err)
	}
	user.RemoveLinkedAccounts(sgc, []models.PolicyAccount{policyAccount})

	fmt.Println(linkedAccount)

}

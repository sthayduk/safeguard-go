package main

import (
	"fmt"
	"log"

	"github.com/sthayduk/safeguard-go/examples/common"
	"github.com/sthayduk/safeguard-go/models"
)

func main() {
	sgc, err := common.InitClient()
	if err != nil {
		log.Fatal(err)
	}

	// Get user with ID 76 ("Stefan Hayduk")
	user, err := models.GetUser(sgc, 76, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Get both policy accounts at once (133: "da-andresen", 134: "sa-andresen")
	accounts := make([]models.PolicyAccount, 0, 2)
	for _, id := range []int{133, 134} {
		account, err := models.GetPolicyAccount(sgc, id, nil)
		if err != nil {
			log.Fatal(err)
		}
		accounts = append(accounts, account)
	}

	// Link all accounts at once
	linkedAccounts, err := models.AddLinkedAccounts(sgc, user, accounts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Successfully linked %d accounts\n", len(linkedAccounts))
}

package main

import (
	"fmt"
	"log"

	. "github.com/sthayduk/safeguard-go"
	"github.com/sthayduk/safeguard-go/examples/common"
)

func main() {
	sgc, err := common.InitClient()
	if err != nil {
		log.Fatal(err)
	}

	// Get user with ID 76 ("Stefan Hayduk")
	user, err := GetUser(sgc, 76, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Get both policy accounts at once (133: "da-andresen", 134: "sa-andresen")
	accounts := make([]PolicyAccount, 0, 2)
	for _, id := range []int{133, 134} {
		account, err := GetPolicyAccount(sgc, id, nil)
		if err != nil {
			log.Fatal(err)
		}
		accounts = append(accounts, account)
	}

	// Unlink all accounts at once
	unlinkedAccounts, err := RemoveLinkedAccounts(sgc, user, accounts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Successfully unlinked %d accounts\n", len(unlinkedAccounts))
}

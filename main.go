package main

import (
	"log"
	"os"

	"github.com/hashgraph/hedera-sdk-go/v2"
	configs "github.com/hnamzian/hedera-example-go/configs"
)

func main() {
	log := log.New(os.Stdout, "hedera", log.LstdFlags)

	cfg := configs.New(log)

	cfg.Load()

	node := make(map[string]hedera.AccountID, 1)
	node[cfg.NetworkNodeAddress] = cfg.NetworkNodeAccountID

	client := hedera.ClientForNetwork(node)
	client.SetMirrorNetwork([]string{cfg.MirrorNodeAddress})

	client.SetOperator(cfg.AccountID, cfg.PrivateKey)

	// make a CreateAccount transaciton by operator account
	tr, err := hedera.NewAccountCreateTransaction().
		SetKey(cfg.PrivateKey).
		SetInitialBalance(hedera.HbarFromTinybar(1000)).
		Execute(client)
	if err != nil {
		log.Panicf("Unable to Create New Account %s", err)
	}

	receipt, err := tr.GetReceipt(client)
	if err != nil {
		log.Panicf("Unable to get transaction receipt %s", err)
	}
	newAccountId := receipt.AccountID
	log.Printf("New Account created %s", newAccountId)
}

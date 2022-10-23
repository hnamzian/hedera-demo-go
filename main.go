package main

import (
	"log"
	"os"

	configs "github.com/hnamzian/hedera-example-go/configs"
	hedera_client "github.com/hnamzian/hedera-example-go/hedera"
)

func main() {
	log := log.New(os.Stdout, "hedera", log.LstdFlags)

	cfg := configs.New(log)

	cfg.Load()

	hederaConfigs := &hedera_client.HederaClientConfigs{
		NetworkNodeAddress:   cfg.NetworkNodeAddress,
		MirrorNodeAddress:    cfg.MirrorNodeAddress,
		NetworkNodeAccountID: cfg.NetworkNodeAccountID,
		OperatorAccountID:    cfg.AccountID,
		OperatorPrivateKey:   cfg.PrivateKey,
	}
	hc := hedera_client.New(log, hederaConfigs)

	// make a CreateAccount transaciton by operator account
	acc, err := hc.NewAccount()	
	if err != nil {
		log.Fatalf("Unable to create new Account %s", err)
	}
	log.Printf("New Account created %s", acc)
}

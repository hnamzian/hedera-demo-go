package configs

import (
	"log"
	"os"

	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/joho/godotenv"
)

type Configs struct {
	log *log.Logger

	AccountID            hedera.AccountID  `json:"accountId"`
	PrivateKey           hedera.PrivateKey `json:"privateKey"`
	NetworkNodeAddress   string            `json:"networkNodeAddress"`
	MirrorNodeAddress    string            `json:"mirrorNodeAddress"`
	NetworkNodeAccountID hedera.AccountID  `json:"networkNodeAccountId"`
}

func New(log *log.Logger) *Configs {
	return &Configs{log, hedera.AccountID{}, hedera.PrivateKey{}, "", "", hedera.AccountID{}}
}

func (cfg *Configs) Load() {
	err := godotenv.Load(".env")
	if err != nil {
		cfg.log.Panicf("Unable to load environment variables %s", err)
	}

	cfg.AccountID, err = hedera.AccountIDFromString(os.Getenv("ACCOUNT_ID"))
	if err != nil {
		cfg.log.Panicf("Unable to load Account ID from ENV %s", err)
	}

	cfg.PrivateKey, err = hedera.PrivateKeyFromString(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		cfg.log.Panicf("Unable to load Private Key from ENV %s", err)
	}

	cfg.NetworkNodeAddress = os.Getenv("NETWORK_NODE_ADDRESS")
	if len(cfg.NetworkNodeAddress) == 0 {
		cfg.log.Panicf("Unable to load Network Node Address from ENV %s", err)
	}

	cfg.MirrorNodeAddress = os.Getenv("MIRROR_NODE_ADDRESS")
	if len(cfg.MirrorNodeAddress) == 0 {
		cfg.log.Panicf("Unable to load Mirror Node Address from ENV %s", err)
	}

	cfg.NetworkNodeAccountID, err = hedera.AccountIDFromString(os.Getenv("NODE_ACCOUNT_ID"))
	if err != nil {
		cfg.log.Panicf("Unable to load Node Account ID from ENV %s", err)
	}

	cfg.log.Printf("Account ID %s", cfg.AccountID)
	cfg.log.Printf("Private Key %s", cfg.PrivateKey)
	cfg.log.Printf("Network Node Address %s", cfg.NetworkNodeAccountID)
	cfg.log.Printf("Mirror Node Address %s", cfg.MirrorNodeAddress)
	cfg.log.Printf("Node Account ID %s", cfg.NetworkNodeAccountID)
}

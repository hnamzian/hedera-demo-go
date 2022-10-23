package hedera_client

import (
	"log"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

type HederaClient struct {
	log     *log.Logger
	client  *hedera.Client
	configs *HederaClientConfigs
}

type HederaClientConfigs struct {
	NetworkNodeAddress   string
	MirrorNodeAddress    string
	NetworkNodeAccountID hedera.AccountID
	OperatorAccountID    hedera.AccountID
	OperatorPrivateKey   hedera.PrivateKey
}

func New(log *log.Logger, configs *HederaClientConfigs) *HederaClient {
	node := make(map[string]hedera.AccountID, 1)
	node[configs.NetworkNodeAddress] = configs.NetworkNodeAccountID

	client := hedera.ClientForNetwork(node)
	client.SetMirrorNetwork([]string{configs.MirrorNodeAddress})

	client.SetOperator(configs.OperatorAccountID, configs.OperatorPrivateKey)

	return &HederaClient{log, client, configs}
}

func (hc *HederaClient) NewAccount() (*hedera.AccountID, error) {
	tr, err := hedera.NewAccountCreateTransaction().
		SetKey(hc.configs.OperatorPrivateKey).
		SetInitialBalance(hedera.HbarFromTinybar(1000)).
		Execute(hc.client)
	if err != nil {
		return nil, err
	}

	receipt, err := tr.GetReceipt(hc.client)
	if err != nil {
		return nil, err
	}

	acc := receipt.AccountID

	return acc, nil
}

package eth

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthClient struct {
	ethclient *ethclient.Client
}

func (e *EthClient) NodeListContract(contractAddress string) *Nodelist {
	nodeListAddress := common.HexToAddress(contractAddress)
	NodeListContract, err := NewNodelist(nodeListAddress, e.ethclient)
	if err != nil {
		panic(err)
	}

	return NodeListContract

}

func NewEthClient(ethURL string) *EthClient {
	client, err := ethclient.Dial(ethURL)
	if err != nil {
		panic(err)
	}
	return &EthClient{
		ethclient: client,
	}
}

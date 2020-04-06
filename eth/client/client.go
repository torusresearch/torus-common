package eth

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/torusresearch/torus-common/eth/gobindings/nodelist"
)

type EthClient struct {
	ethclient *ethclient.Client
}

func (e *EthClient) Client() *ethclient.Client {
	return e.ethclient
}

func (e *EthClient) NodeListContract(nodeListAddress common.Address) *nodelist.NodeList {
	NodeListContract, err := nodelist.NewNodeList(nodeListAddress, e.ethclient)
	if err != nil {
		panic(err)
	}

	return NodeListContract

}

func (e *EthClient) NodeListContractFromAddressString(contractAddress string) *nodelist.NodeList {
	nodeListAddress := common.HexToAddress(contractAddress)
	return e.NodeListContract(nodeListAddress)

}

// DeployContract is blocking, and depending on what network you use can take
// anything from a second (ganache) to 2 minutes (mainnet)
func (e *EthClient) DeployContract(key *ecdsa.PrivateKey) (common.Address, *types.Transaction, *nodelist.NodeList, error) {
	blockchain := e.ethclient

	// Get credentials for the account to charge for contract deployments
	auth := bind.NewKeyedTransactor(key)

	// address, _, _, error := DeployNodeList(
	// 	auth,
	// 	blockchain,
	// )

	return nodelist.DeployNodeList(
		auth,
		blockchain,
	)

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

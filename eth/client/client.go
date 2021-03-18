package eth

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/torusresearch/torus-common/eth/gobindings/nodeListProxy"
	"github.com/torusresearch/torus-common/eth/gobindings/nodelist"
	"github.com/torusresearch/torus-common/eth/gobindings/verifierListAggregator"
	"github.com/torusresearch/torus-common/eth/gobindings/verifierOwnerManager"
	"github.com/torusresearch/torus-common/eth/gobindings/verifierlist"
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
func (e *EthClient) DeployContract(key *ecdsa.PrivateKey) (nodeListAddress common.Address, nodeListContract *nodelist.NodeList, verifierListAddress common.Address, verifierListContract *verifierlist.VerifierList, err error) {
	blockchain := e.ethclient

	// Get credentials for the account to charge for contract deployments
	auth := bind.NewKeyedTransactor(key)

	nodeListAddress, _, nodeListContract, err = nodelist.DeployNodeList(
		auth,
		blockchain,
	)
	if err != nil {
		return
	}

	verifierListAddress, _, verifierListContract, err = verifierlist.DeployVerifierList(auth, blockchain)
	if err != nil {
		return
	}

	return
}

// DeployAdditionalContracts deploys additional contracts to simulate the production setup
func (e *EthClient) DeployAdditionalContracts(key *ecdsa.PrivateKey, verifierListAddress, nodeListAddress common.Address, epoch *big.Int) (verifierOwnerManagerAddress common.Address, verifierOwnerManagerContract *verifierOwnerManager.VerifierOwnerManager, verifierListAggregatorAddress common.Address, verifierListAggregatorContract *verifierListAggregator.VerifierListAggregator, nodeListProxyAddress common.Address, nodeListProxyContract *nodeListProxy.NodeListProxy, err error) {
	blockchain := e.ethclient

	// Get credentials for the account to charge for contract deployments
	auth := bind.NewKeyedTransactor(key)

	verifierOwnerManagerAddress, _, verifierOwnerManagerContract, err = verifierOwnerManager.DeployVerifierOwnerManager(
		auth,
		blockchain,
		verifierListAddress,
	)
	if err != nil {
		return
	}

	verifierListAggregatorAddress, _, verifierListAggregatorContract, err = verifierListAggregator.DeployVerifierListAggregator(auth, blockchain, verifierListAddress, verifierOwnerManagerAddress)
	if err != nil {
		return
	}

	nodeListProxyAddress, _, nodeListProxyContract, err = nodeListProxy.DeployNodeListProxy(auth, blockchain, nodeListAddress, epoch)
	if err != nil {
		return
	}

	return
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

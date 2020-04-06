package eth

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/torusresearch/torus-common/eth/gobindings/nodelist"
)

// Test inbox contract gets deployed correctly
func TestDeployNodelist(t *testing.T) {

	//Setup simulated block chain
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000)}
	blockchain := backends.NewSimulatedBackend(alloc, 0)

	//Deploy contract
	address, _, nodelistInstance, err := nodelist.DeployNodeList(
		auth,
		blockchain,
	)
	// commit all pending transactions
	blockchain.Commit()

	fmt.Println("contract deployed @", address.Hex())
	if err != nil {
		t.Fatalf("Failed to deploy the NodeList contract: %v", err)
	}

	if len(address.Bytes()) == 0 {
		t.Error("Expected a valid deployment address. Received empty address byte array instead")
	}

	_, err = nodelist.EpochInfo(nil, big.NewInt(0))

	if err != nil {
		t.Fatal(err)
	}

	// Epoch
	_, err = nodelistInstance.UpdateEpoch(auth, big.NewInt(1), big.NewInt(5), big.NewInt(3), big.NewInt(1), []common.Address{}, big.NewInt(0), big.NewInt(2))
	if err != nil {
		t.Fatal(err)
	}
	blockchain.Commit()

	// Accounts to whitelist
	whitelistedAccounts := []string{"0x52c476751142ce2fb4db4f19b500e78feee10b06", "0xff364b6b86ea5a4f59cc4989da23b833dac15304", "0xdc0dd04aac998e8aa9f2de236b3ba04ddafd26ca", "0x253db77f1ae216722b2f67f33ef3c8e00b2689e6", "0x271346169993368f94cb2c443b8b8cdbdd5edf04", "0xa0ae28ec27fea7a577b21330f6ce8ae45a55fe76", "0xf34a875cffe643d44546b76f0c9412dfb9d2b379", "0x35d946c9c4598cd2eaee5754ce2041911dc816ce", "0xd6ee5e06ac11a62fd0be1912debeeb4abc24f723", "0x40fa4b9e4411e7f5f58713eff426cad4f0294ab5", "0x0cda757357158e4d8ad94433e36f1fe05a1dc576", "0xa22e3c16264dc688107142776139d1fb4bb9d549", "0x0b998b7229bfd254acf50b4e2739e73d937dc1c9", "0xfc54c26e24b4570590c11486bd627aa4b7339523", "0xb572081928b988abe713ffe60f8cf28ef80eee07", "0xd54e0c310a97916e67d07aa501f74524e82c3af1", "0xaba31e255b490365584a56f4ebc5037963e584d5", "0x3ecefafea7db9d0e26dc0d266504587cb66f6008", "0x184b56d50300b4cd604a587491cb7bcb0ffc7454", "0xd6eca392ada22e18c9eebde2828b38e66813af5f"}
	for _, acc := range whitelistedAccounts {
		accAddress := common.HexToAddress(acc)
		_, err := nodelistInstance.UpdateWhitelist(auth, big.NewInt(1), accAddress, true)
		if err != nil {
			t.Fatal(err)
		}
		blockchain.Commit()

		isWhitelisted, err := nodelistInstance.IsWhitelisted(nil, big.NewInt(1), accAddress)
		if err != nil {
			t.Fatal(err)
		}

		if !isWhitelisted {
			t.Fatal("account", acc, "expected to be whitelisted, but it is not!")
		}
	}

}

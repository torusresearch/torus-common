package eth

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
)

const testPrivKey = "2cc22a5c9b07f1286d7e708ea326b4a0ac18a70decfc320176cf5f103a5732ca"

// TestGanacheDeploy depends on the local installation of ganache
func TestGanacheDeploy(t *testing.T) {
	c := NewEthClient("http://localhost:8545/")
	key, err := crypto.HexToECDSA(testPrivKey)
	if err != nil {
		t.Fatal(err)
	}

	address, tx, _, err := c.DeployContract(key)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("The address of deployed contract is: %s \n", address.Hex())
	fmt.Println(tx.Hash().Hex())
}

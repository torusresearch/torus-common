package secp256k1

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/torusresearch/torus-common/common"
)

func TestEncryptDecrypt(test *testing.T) {
	input := "hello"
	privK := *big.NewInt(1234)
	pubKx, pubKy := Curve.ScalarBaseMult(privK.Bytes())
	pubK := common.Point{X: *pubKx, Y: *pubKy}
	encryptedOutput, err := Encrypt(pubK, []byte(input))
	if err != nil {
		test.Fail()
	}
	decryption, err := Decrypt(privK, encryptedOutput)
	if err != nil {
		test.Fail()
	}
	assert.Equal(test, string(decryption), input)
}

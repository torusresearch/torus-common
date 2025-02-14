package secp256k1

import (
	"bytes"
	"fmt"
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

	fmt.Printf("Private Key: %x\n", privK.Bytes())
	fmt.Printf("Public Key X: %x\n", pubKx.Bytes())
	fmt.Printf("Public Key Y: %x\n", pubKy.Bytes())

	encryptedOutput, err := Encrypt(pubK, []byte(input))
	if err != nil {
		test.Fatalf("Encryption failed: %v", err)
	}

	fmt.Printf("Encrypted output length: %d\n", len(encryptedOutput))
	fmt.Printf("Encrypted output: %x\n", encryptedOutput)

	decryption, err := Decrypt(privK, encryptedOutput)
	if err != nil {
		test.Fatalf("Decryption failed: %v", err)
	}

	fmt.Printf("Input: %s | Decrypted: %s\n", input, string(decryption))

	assert.Equal(test, string(decryption), input)
}

func TestEncryptDecryptMore(t *testing.T) {
	// Test cases
	testCases := []struct {
		name  string
		input []byte
	}{
		{"Empty input", []byte{}},
		{"Hello World", []byte("Hello World")},
		{"Binary data", []byte{0x00, 0x01, 0x02, 0x03, 0xFF}},
		{"Long text", []byte("This is a longer text to test encryption and decryption of larger data")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Generate a random private key
			privKey := RandomBigInt()

			// Calculate public key
			x, y := Curve.ScalarBaseMult(privKey.Bytes())
			pubKey := common.Point{X: *x, Y: *y}

			// Encrypt
			encrypted, err := Encrypt(pubKey, tc.input)
			if err != nil {
				t.Fatalf("Encryption failed: %v", err)
			}

			fmt.Printf("Encrypted output length: %d of input: %s\n", len(encrypted), string(tc.input))

			// Decrypt
			decrypted, err := Decrypt(*privKey, encrypted)
			if err != nil {
				t.Fatalf("Decryption failed: %v", err)
			}

			// Compare original and decrypted data
			if !bytes.Equal(tc.input, decrypted) {
				t.Errorf("Decrypted data doesn't match original input.\nOriginal: %x\nDecrypted: %x",
					tc.input, decrypted)
			}
		})
	}
}

// Test that different encryptions of the same message are different
func TestEncryptionRandomness(t *testing.T) {
	input := []byte("Hello World")
	privKey := RandomBigInt()
	x, y := Curve.ScalarBaseMult(privKey.Bytes())
	pubKey := common.Point{X: *x, Y: *y}

	encrypted1, err := Encrypt(pubKey, input)
	if err != nil {
		t.Fatalf("First encryption failed: %v", err)
	}

	encrypted2, err := Encrypt(pubKey, input)
	if err != nil {
		t.Fatalf("Second encryption failed: %v", err)
	}

	if bytes.Equal(encrypted1, encrypted2) {
		t.Error("Two encryptions of the same message should be different")
	}
}

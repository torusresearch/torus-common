package crytpo

import (
	"crypto/ecdsa"
	"encoding/hex"

	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/prometheus/common/log"
	"github.com/torusresearch/bijson"
	"github.com/torusresearch/torus-common/secp256k1"
)

type SignObject struct {
	Body    []byte `json:"body"`
	PubKeyX string `json:"pubkeyx"`
	PubKeyY string `json:"pubkeyy"`
}

// SignData returns a hex-encoded signature of the passed in data.
// NOTE: this function has basically been copied over from `utils.go` in
// torus-public, however we only deal with hex encoded signatures
func SignData(rawData []byte, privateKey *ecdsa.PrivateKey) string {
	hashRaw := secp256k1.Keccak256(rawData)
	signature, err := ethCrypto.Sign(hashRaw, privateKey)
	if err != nil {
		log.Fatal(err)
	}

	return hex.EncodeToString(signature)
}

func SignDataWithPubKey(bodyBytes []byte, pubKeyX string, pubKeyY string, privateKey *ecdsa.PrivateKey) string {
	objectToSign := SignObject{
		bodyBytes,
		pubKeyX,
		pubKeyY,
	}

	rawData, err := bijson.Marshal(objectToSign)
	if err != nil {
		log.Fatal(err)
	}

	return SignData(rawData, privateKey)
}

func verify(pubkey, hash, signature []byte) bool {
	return ethCrypto.VerifySignature(pubkey, hash, signature)
}

// Verify checks whether the signature matches the data
func Verify(encodedPubKey, signature string, data []byte) bool {
	pubKey, err := hex.DecodeString(encodedPubKey)
	if err != nil {
		log.Error(err)
		return false
	}

	hashRaw := secp256k1.Keccak256(data)
	sig, err := hex.DecodeString(signature)
	if err != nil {
		log.Error(err)
		return false
	}

	return ethCrypto.VerifySignature(pubKey, hashRaw, sig)
}

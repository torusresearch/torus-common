package secp256k1

import (
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/torusresearch/torus-common/common"
	"golang.org/x/crypto/sha3"
)

type KoblitzCurve struct {
	*btcec.KoblitzCurve
}

var (
	Curve = &KoblitzCurve{btcec.S256()}
	// field order, also known as p, usually used for scalars
	FieldOrder = common.HexToBigInt("fffffffffffffffffffffffffffffffffffffffffffffffffffffffefffffc2f")
	// group order, also known as q, it is the number of points in the curve, and is usually used in exponents
	GeneratorOrder = common.HexToBigInt("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141")
	// scalar to the power of this is like square root, eg. y^sqRoot = y^0.5 (if it exists)
	SqRoot = common.HexToBigInt("3fffffffffffffffffffffffffffffffffffffffffffffffffffffffbfffff0c")
	G      = common.Point{X: *Curve.Gx, Y: *Curve.Gy}
	H      = *HashToPoint(G.X.Bytes())
)

func RandomBigInt() *big.Int {
	randomInt, _ := rand.Int(rand.Reader, GeneratorOrder)
	return randomInt
}

func HashToPoint(data []byte) *common.Point {
	keccakHash := Keccak256(data)
	x := new(big.Int)
	x.SetBytes(keccakHash)
	for {
		beta := new(big.Int)
		beta.Exp(x, big.NewInt(3), FieldOrder)
		beta.Add(beta, big.NewInt(7))
		beta.Mod(beta, FieldOrder)
		y := new(big.Int)
		y.Exp(beta, SqRoot, FieldOrder)
		if new(big.Int).Exp(y, big.NewInt(2), FieldOrder).Cmp(beta) == 0 {
			return &common.Point{X: *x, Y: *y}
		} else {
			x.Add(x, big.NewInt(1))
		}
	}
}

func Keccak256(data ...[]byte) []byte {
	d := sha3.NewLegacyKeccak256()
	for _, b := range data {
		n, err := d.Write(b)
		if err != nil {
			fmt.Printf("Could not write n %v err %v", n, err)
		}
	}
	return d.Sum(nil)
}

func Encrypt(pubKey common.Point, input []byte) (encryptedOutput []byte, err error) {
	// Convert to ECDSA public key first
	ecdsaPub := &ecdsa.PublicKey{
		Curve: btcec.S256(),
		X:     &pubKey.X,
		Y:     &pubKey.Y,
	}

	// Convert to ecies public key and set parameters
	pub := ecies.ImportECDSAPublic(ecdsaPub)
	pub.Params = ecies.ECIES_AES128_SHA256 // Use standard parameters

	// Encrypt using ECIES
	return ecies.Encrypt(rand.Reader, pub, input, nil, nil)
}

func Decrypt(privKey big.Int, input []byte) (decryptedOutput []byte, err error) {
	// Convert to ECDSA private key first
	ecdsaPriv := &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: btcec.S256(),
		},
		D: &privKey,
	}

	// Calculate public key components
	ecdsaPriv.PublicKey.X, ecdsaPriv.PublicKey.Y = btcec.S256().ScalarBaseMult(privKey.Bytes())

	// Convert to ecies private key and set parameters
	priv := ecies.ImportECDSA(ecdsaPriv)
	priv.Params = ecies.ECIES_AES128_SHA256 // Use standard parameters

	// Decrypt using ECIES
	return priv.Decrypt(input, nil, nil)
}

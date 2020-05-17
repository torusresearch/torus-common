package secp256k1

import (
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/btcsuite/btcd/btcec"
	"github.com/torusresearch/torus-common/common"
	"golang.org/x/crypto/sha3"
)

type KoblitzCurve struct {
	*btcec.KoblitzCurve
}

var (
	Curve *KoblitzCurve
	// field order, also known as p, usually used for scalars
	FieldOrder = common.HexToBigInt("fffffffffffffffffffffffffffffffffffffffffffffffffffffffefffffc2f")
	// group order, also known as q, it is the number of points in the curve, and is usually used in exponents
	GeneratorOrder = common.HexToBigInt("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141")
	// scalar to the power of this is like square root, eg. y^sqRoot = y^0.5 (if it exists)
	SqRoot = common.HexToBigInt("3fffffffffffffffffffffffffffffffffffffffffffffffffffffffbfffff0c")
	G      common.Point
	H      common.Point
)

func init() {
	P, _ := new(big.Int).SetString("1000000000000000000000000000000014def9dea2f79cd65812631a5cf5d3ed", 16)
	N, _ := new(big.Int).SetString("254", 10)
	B, _ := new(big.Int).SetString("14def9dea2f79cd65812631a5cf5d3ed", 16)
	Gx, _ := new(big.Int).SetString("216936D3CD6E53FEC0A4E231FDD6DC5C692CC7609525A7B2C9562D608F25D51A", 16)
	Gy, _ := new(big.Int).SetString("6666666666666666666666666666666666666666666666666666666666666658", 16)

	Curve = &KoblitzCurve{
		KoblitzCurve: &btcec.KoblitzCurve{
			CurveParams: &elliptic.CurveParams{
				P:       P,
				N:       N,
				B:       B,
				Gx:      Gx,
				Gy:      Gy,
				BitSize: 256,
				Name:    "ed25519",
			},
		},
	}

	G = common.Point{X: *Gx, Y: *Gy}
	H = *HashToPoint(G.X.Bytes())
}

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
	return btcec.Encrypt(&btcec.PublicKey{
		Curve: Curve,
		X:     &pubKey.X,
		Y:     &pubKey.Y,
	}, input)
}

func Decrypt(privKey big.Int, input []byte) (decryptedOutput []byte, err error) {
	return btcec.Decrypt(&btcec.PrivateKey{
		D: &privKey,
	}, input)
}

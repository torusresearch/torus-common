package secp256k1

import (
	"crypto/rand"
	"math/big"

	secp256k1 "github.com/btcsuite/btcd/btcec"
	"golang.org/x/crypto/sha3"
)

type KoblitzCurve struct {
	*secp256k1.KoblitzCurve
}

type Point struct {
	X big.Int
	Y big.Int
}

var (
	Curve = &KoblitzCurve{secp256k1.S256()}
	// field order, also known as p, usually used for scalars
	FieldOrder = HexToBigInt("fffffffffffffffffffffffffffffffffffffffffffffffffffffffefffffc2f")
	// group order, also known as q, it is the number of points in the curve, and is usually used in exponents
	GeneratorOrder = HexToBigInt("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141")
	// scalar to the power of this is like square root, eg. y^sqRoot = y^0.5 (if it exists)
	SqRoot = HexToBigInt("3fffffffffffffffffffffffffffffffffffffffffffffffffffffffbfffff0c")
	G      = Point{X: *Curve.Gx, Y: *Curve.Gy}
	H      = *HashToPoint(G.X.Bytes())
)

func HexToBigInt(s string) *big.Int {
	r, ok := new(big.Int).SetString(s, 16)
	if !ok {
		panic("invalid hex in source file: " + s)
	}
	return r
}

func RandomBigInt() *big.Int {
	randomInt, _ := rand.Int(rand.Reader, GeneratorOrder)
	return randomInt
}
func HashToPoint(data []byte) *Point {
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
			return &Point{X: *x, Y: *y}
		} else {
			x.Add(x, big.NewInt(1))
		}
	}
}

func Keccak256(data ...[]byte) []byte {
	d := sha3.NewLegacyKeccak256()
	for _, b := range data {
		d.Write(b)
	}
	return d.Sum(nil)
}

package lib

import (
	"crypto/elliptic"
	"math/big"
	"sync"
)

var ed25519Initonce sync.Once
var ed25519 Ed25519Curve

type Ed25519Curve struct {
	*elliptic.CurveParams
}

func ed25519InitAll() {
	g := ED25519().NewGeneratorPoint()
	gx, gy := coordsFromPoint(g)
	// taken from https://datatracker.ietf.org/doc/html/rfc8032
	ed25519.CurveParams = new(elliptic.CurveParams)
	ed25519.P, _ = new(big.Int).SetString("7FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFED", 16)
	ed25519.N = generatorOrder
	ed25519.Gx = gx
	ed25519.Gy = gy
	ed25519.BitSize = 255
	ed25519.Name = "ed25519"
}

func coordsFromPoint(p Point) (*big.Int, *big.Int) {
	au := p.ToAffineUncompressed()
	return newBigIntLE(au[:32]), newBigIntLE(au[32:])
}

// newBigIntLE creates a new big.Int from little endian bytes.
func newBigIntLE(a []byte) *big.Int {
	b := make([]byte, len(a))
	for i, a_i := range a {
		b[len(b)-1-i] = a_i
	}
	return new(big.Int).SetBytes(b)
}

func Ed25519() *Ed25519Curve {
	ed25519Initonce.Do(ed25519InitAll)
	return &ed25519
}

func (curve *Ed25519Curve) Params() *elliptic.CurveParams {
	return curve.CurveParams
}

func (curve *Ed25519Curve) IsOnCurve(p Point) bool {
	return p.IsOnCurve()
}

func (curve *Ed25519Curve) Add(p1, p2 Point) Point {
	return p1.Add(p2)
}

func (curve *Ed25519Curve) Double(p Point) Point {
	return p.Double()
}

func (curve *Ed25519Curve) ScalarMult(p Point, s Scalar) Point {
	return p.Mul(s)
}

func (curve *Ed25519Curve) ScalarBaseMult(s Scalar) Point {
	return ED25519().ScalarBaseMult(s)
}

func (curve *Ed25519Curve) Neg(p Point) Point {
	return p.Neg()
}

func (curve *Ed25519Curve) Hash(msg []byte) Point {
	return ED25519().Point.Hash(msg)
}

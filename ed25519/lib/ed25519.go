package lib

import (
	"crypto/elliptic"
	"math/big"
	"sync"

	"github.com/bwesterb/go-ristretto/edwards25519"
)

var ed25519Initonce sync.Once
var ed25519 Ed25519Curve

type Ed25519Curve struct {
	*elliptic.CurveParams
}

func ed25519InitAll() {
	generatorX, generatorY := getGenerator()
	// taken from https://datatracker.ietf.org/doc/html/rfc8032
	ed25519.CurveParams = new(elliptic.CurveParams)
	ed25519.P, _ = new(big.Int).SetString("7FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFED", 16)
	ed25519.N = generatorOrder
	ed25519.Gx = generatorX
	ed25519.Gy = generatorY
	ed25519.BitSize = 255
	ed25519.Name = "ed25519"
}

func getGenerator() (*big.Int, *big.Int) {
	generator := ED25519().NewGeneratorPoint()
	generatorBytes := generator.ToAffineUncompressed()
	var genX, genY [32]byte
	copy(genX[:], generatorBytes[:32])
	copy(genY[:], generatorBytes[32:])

	var x, y edwards25519.FieldElement
	x.SetBytes(&genX)
	y.SetBytes(&genY)

	return x.BigInt(), y.BigInt()
}

func Ed25519() *Ed25519Curve {
	ed25519Initonce.Do(ed25519InitAll)
	return &ed25519
}

func (curve *Ed25519Curve) Params() *elliptic.CurveParams {
	return curve.CurveParams
}

func (curve *Ed25519Curve) IsOnCurve(x, y *big.Int) bool {
	ed25519Curve := ED25519()
	p, err := ed25519Curve.Point.Set(x, y)
	if err != nil {
		return false
	}
	return p.IsOnCurve()
}

func (curve *Ed25519Curve) Add(x1, y1, x2, y2 *big.Int) (*big.Int, *big.Int, error) {
	ed25519Curve := ED25519()
	p1, err := ed25519Curve.Point.Set(x1, y1)
	if err != nil {
		return nil, nil, err
	}
	p2, err := ed25519Curve.Point.Set(x2, y2)
	if err != nil {
		return nil, nil, err
	}

	resBytes := p1.Add(p2).ToAffineUncompressed()
	var xBytes, yBytes [32]byte
	copy(xBytes[:], resBytes[:32])
	copy(yBytes[:], resBytes[32:])

	var x, y edwards25519.FieldElement
	x.SetBytes(&xBytes)
	y.SetBytes(&yBytes)

	return x.BigInt(), y.BigInt(), nil
}

func (curve *Ed25519Curve) Double(x1, y1 *big.Int) (*big.Int, *big.Int, error) {
	ed25519Curve := ED25519()
	p, err := ed25519Curve.Point.Set(x1, y1)
	if err != nil {
		return nil, nil, err
	}

	resBytes := p.Double().ToAffineUncompressed()
	var xBytes, yBytes [32]byte
	copy(xBytes[:], resBytes[:32])
	copy(yBytes[:], resBytes[32:])

	var x, y edwards25519.FieldElement
	x.SetBytes(&xBytes)
	y.SetBytes(&yBytes)

	return x.BigInt(), y.BigInt(), nil
}

func (curve *Ed25519Curve) ScalarMult(Bx, By, k *big.Int) (*big.Int, *big.Int, error) {
	ed25519Curve := ED25519()
	p, err := ed25519Curve.Point.Set(Bx, By)
	if err != nil {
		return nil, nil, err
	}

	s, err := ED25519().Scalar.SetBigInt(k)
	if err != nil {
		return nil, nil, err
	}

	resBytes := p.Mul(s).ToAffineUncompressed()
	var xBytes, yBytes [32]byte
	copy(xBytes[:], resBytes[:32])
	copy(yBytes[:], resBytes[32:])

	var x, y edwards25519.FieldElement
	x.SetBytes(&xBytes)
	y.SetBytes(&yBytes)

	return x.BigInt(), y.BigInt(), nil
}

func (curve *Ed25519Curve) ScalarBaseMult(k *big.Int) (*big.Int, *big.Int, error) {
	s, err := ED25519().Scalar.SetBigInt(k)
	if err != nil {
		return nil, nil, err
	}
	resBytes := ED25519().ScalarBaseMult(s).ToAffineUncompressed()
	var xBytes, yBytes [32]byte
	copy(xBytes[:], resBytes[:32])
	copy(yBytes[:], resBytes[32:])

	var x, y edwards25519.FieldElement
	x.SetBytes(&xBytes)
	y.SetBytes(&yBytes)

	return x.BigInt(), y.BigInt(), nil
}

func (curve *Ed25519Curve) Neg(Bx, By *big.Int) (*big.Int, *big.Int, error) {
	ed25519Curve := ED25519()
	p, err := ed25519Curve.Point.Set(Bx, By)
	if err != nil {
		return nil, nil, err
	}

	resBytes := p.Neg().ToAffineUncompressed()
	var xBytes, yBytes [32]byte
	copy(xBytes[:], resBytes[:32])
	copy(yBytes[:], resBytes[32:])

	var x, y edwards25519.FieldElement
	x.SetBytes(&xBytes)
	y.SetBytes(&yBytes)

	return x.BigInt(), y.BigInt(), nil
}

func (curve *Ed25519Curve) Hash(msg []byte) (*big.Int, *big.Int) {
	ed25519Curve := ED25519()
	data := ed25519Curve.Point.Hash(msg).ToAffineUncompressed()
	var xBytes, yBytes [32]byte
	copy(xBytes[:], data[:32])
	copy(yBytes[:], data[32:])
	var x, y edwards25519.FieldElement
	x.SetBytes(&xBytes)
	y.SetBytes(&yBytes)
	return x.BigInt(), y.BigInt()
}

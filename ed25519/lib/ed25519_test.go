package lib

import (
	"encoding/hex"
	ed "filippo.io/edwards25519"
	"github.com/bwesterb/go-ristretto/edwards25519"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHash(t *testing.T) {
	curve := ED25519()
	ed25519 := Ed25519()
	sc := curve.Point.Hash(ed25519.Gx.Bytes())
	s, ok := sc.(*PointEd25519)
	assert.True(t, ok)

	assert.True(t, s.IsOnCurve())

	data := s.ToAffineUncompressed()
	var xBytes, yBytes [32]byte
	copy(xBytes[:], data[:32])
	copy(yBytes[:], data[32:])
	var x, y edwards25519.FieldElement
	x.SetBytes(&xBytes)
	y.SetBytes(&yBytes)

	p, err := curve.Point.Set(x.BigInt(), y.BigInt())
	assert.Nil(t, err)

	assert.True(t, p.IsOnCurve())
}

func toRPt(hx string) *ed.Point {
	e, _ := hex.DecodeString(hx)
	var data [32]byte
	copy(data[:], e)
	pt, _ := new(PointEd25519).FromAffineCompressed(data[:])
	return pt.(*PointEd25519).Value
}

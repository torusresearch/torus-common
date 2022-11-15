package ed25519

import (
	"fmt"
	ed25519 "github.com/coinbase/kryptology/pkg/sharing/v1"
	"github.com/torusresearch/torus-common/common"
	"golang.org/x/crypto/sha3"
)

type EdwardsCurve struct {
	*ed25519.Ed25519Curve
}

// GeneratorOrder = 2^252 + 27742317777372353535851937790883648493 = 1000000000000000000000000000000014DEF9DEA2F79CD65812631A5CF5D3ED = Curve.N
// FieldOrder = 2^255 - 19 = 7FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFED = Curve.P

var (
	Curve = &EdwardsCurve{ed25519.Ed25519()}
	G     = common.Point{X: *Curve.Gx, Y: *Curve.Gy}
	H     = HashToPoint(G.X.Bytes())
)

func HashToPoint(data []byte) common.Point {
	return common.BigIntToPoint(Curve.Hash(data))
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

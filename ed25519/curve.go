package ed25519

import (
	"fmt"
	"math/big"

	"github.com/torusresearch/torus-common/ed25519/lib"
	"golang.org/x/crypto/sha3"
)

type EdwardsCurve struct {
	*lib.Ed25519Curve
}

// GeneratorOrder = 2^252 + 27742317777372353535851937790883648493 = 1000000000000000000000000000000014DEF9DEA2F79CD65812631A5CF5D3ED = Curve.N
// FieldOrder = 2^255 - 19 = 7FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFED = Curve.P

var (
	Curve = &EdwardsCurve{lib.Ed25519()}
	G     = lib.ED25519().NewGeneratorPoint()
	H     = HashToPoint(G.ToAffineCompressed())
)

func HashToPoint(data []byte) lib.Point {
	return Curve.Hash(data)
}

func HashToScalar(data []byte) lib.Scalar {
	return lib.ED25519().Scalar.Hash(data)
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

func NewScalar(s *big.Int) (lib.Scalar, error) {
	return lib.ED25519().Scalar.SetBigInt(s)
}

type Scalar = lib.Scalar
type Point = lib.Point

var ZeroScalar = lib.ED25519().Scalar.Zero()

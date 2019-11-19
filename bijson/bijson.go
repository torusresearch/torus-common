package bijson

import (
	"fmt"
	"io"
	"math/big"
	"reflect"

	"github.com/ugorji/go/codec"
)

type RawMessage codec.Raw

type BigIntExt struct{}

func (x BigIntExt) ConvertExt(v interface{}) interface{} {
	v2 := v.(*big.Int)
	return v2.Text(16)
}

func (x BigIntExt) UpdateExt(dst interface{}, src interface{}) {
	tt := dst.(*big.Int)
	tt.SetString(src.(string), 16)
}

func (x BigIntExt) WriteExt(v interface{}) []byte {
	v2 := v.(*big.Int)
	return []byte(v2.Text(16))
}

func (x BigIntExt) ReadExt(dest interface{}, v []byte) {
	tt := dest.(*big.Int)
	tt.SetString(string(v), 16)
}

var h codec.CborHandle

// var h codec.MsgpackHandle

var jsonHandle codec.JsonHandle

func init() {
	// h.SetBytesExt(reflect.TypeOf(new(big.Int)), 1, BigIntExt{})
	h.SetInterfaceExt(reflect.TypeOf(new(big.Int)), 1, BigIntExt{})
}

func marshal(v interface{}) ([]byte, error) {
	b := make([]byte, 0, 64)
	enc := codec.NewEncoderBytes(&b, &h)
	if err := enc.Encode(v); err != nil {
		return nil, err
	}

	return b, nil
}

func Unmarshal(data []byte, v interface{}) error {
	dec := codec.NewDecoderBytes(data, &h)
	return dec.Decode(v)
}

func Marshal(v interface{}) ([]byte, error) {
	b, err := marshal(v)
	if err != nil {
		return nil, err
	}

	u := v
	err = Unmarshal(b, &u)
	if !reflect.DeepEqual(u, v) {
		fmt.Println(err)
		fmt.Println(u)
		fmt.Println(v)
		panic(nil)
	}

	return b, nil
}

func NewDecoder(r io.Reader) *codec.Decoder {
	return codec.NewDecoder(r, &h)
}

func NewEncoder(w io.Writer) *codec.Encoder {
	return codec.NewEncoder(w, &h)
}

func JsonMarshal(v interface{}) ([]byte, error) {
	b := make([]byte, 0, 64)
	enc := codec.NewEncoderBytes(&b, &jsonHandle)
	if err := enc.Encode(v); err != nil {
		return nil, err
	}
	return b, nil
}

func JsonUnmarshal(data []byte, v interface{}) error {
	dec := codec.NewDecoderBytes(data, &jsonHandle)
	return dec.Decode(v)
}

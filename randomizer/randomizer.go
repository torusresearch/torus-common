package randomizer

import (
	cryptoRand "crypto/rand"
	"errors"
	"math/big"
	"math/rand"
	"reflect"
	"time"

	"github.com/torusresearch/torus-common/secp256k1"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomBigInt() *big.Int {
	randomInt, _ := cryptoRand.Int(cryptoRand.Reader, secp256k1.GeneratorOrder)
	return randomInt
}

// RandLetterString Generates a random string of the given length
func RandLetterString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// Randomize ...
// Randomizes the values in the given interface
// The value needs to be passed by reference
// Supported Types:
// 		1. Struct
// 		2. Array/Slice of Struct
// 		3. Float64
// 		4. Int
// 		5. Map
// 		6. String
// Supported Types in Struct
// 		1. Int
// 		2. Float64
// 		3. String
// 		4. Map
// Map Support Definition
// 		1. The key should be string
// 		2. Supported value types are: int64, float64, string, scalar struct
func Randomize(value interface{}) error {
	if reflect.TypeOf(value).Kind() != reflect.Ptr {
		return errors.New("not a pointer")
	}

	mutable := reflect.ValueOf(value)
	switch mutable.Elem().Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < mutable.Elem().Len(); i++ {
			randomValueGenerator(mutable.Field(i))
		}

	case reflect.Map:
		// For each value in the map
		// Check the Kind() of the Value
		// and set the random value
		// for Kind() = Struct run it through randomValueGenerator
		for _, e := range mutable.Elem().MapKeys() {
			v := mutable.Elem().MapIndex(e)
			switch v.Kind() {
			case reflect.Int:
				v.SetMapIndex(e, reflect.ValueOf(rand.Int()))
			case reflect.Float64:
				v.SetMapIndex(e, reflect.ValueOf(rand.Float64()))
			case reflect.String:
				v.SetMapIndex(e, reflect.ValueOf(RandLetterString(len(v.String()))))
			case reflect.Struct:
				randomValueGenerator(v)
			}
		}

	case reflect.Struct:
		randomValueGenerator(mutable.Elem())
	}

	return nil
}

func randomValueGenerator(mutable reflect.Value) {
	// For every field in the given reflect.Value
	// Determine the Kind()
	// 	For reflect.Int, assign a random Int
	// 	For reflect.Float64, assign a random float64
	// 	For reflect.String, assign a random String
	// 	For reflect.Struct, run it through the randomValueGenerator
	// 	For reflect.Map,
	// 		For each value in the map
	// 		Check the Kind() of the Value
	// 		and set the random value
	// 		for Kind() = Struct run it through randomValueGenerator
	reflectBigInt := reflect.ValueOf(*new(big.Int)).Type()
	for i := 0; i < mutable.NumField(); i++ {
		if mutable.Field(i).Type() == reflectBigInt {
			mutable.Field(i).Set(reflect.ValueOf(*RandomBigInt()))
		} else {
			switch mutable.Field(i).Kind() {
			case reflect.Int:
				mutable.Field(i).SetInt(int64(rand.Int()))

			case reflect.Float64:
				mutable.Field(i).SetFloat(rand.Float64())

			case reflect.String:
				mutable.Field(i).SetString(RandLetterString(len(mutable.Field(i).String())))

			case reflect.Array, reflect.Slice:
				for j := 0; j < mutable.Field(i).Len(); j++ {
					randomValueGenerator(mutable.Field(i).Index(j))
				}

			case reflect.Map:
				for _, e := range mutable.Field(i).MapKeys() {
					v := mutable.Field(i).MapIndex(e)
					switch v.Elem().Kind() {
					case reflect.Int:
						mutable.Field(i).SetMapIndex(e, reflect.ValueOf(rand.Int()))
					case reflect.String:
						mutable.Field(i).SetMapIndex(e, reflect.ValueOf(RandLetterString(len(v.String()))))
					case reflect.Struct:
						randomValueGenerator(v)
					case reflect.Float64:
						mutable.Field(i).SetMapIndex(e, reflect.ValueOf(rand.Float64()))
					}
				}

			case reflect.Struct:
				randomValueGenerator(mutable.Field(i))
			}
		}
	}
}

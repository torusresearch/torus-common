package randomizer

import (
	"errors"
	"math/rand"
	"reflect"
)

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

	// Get the mutable reflect reference to the interface{}
	mutable := reflect.ValueOf(value)

	// Check the Kind()
	switch mutable.Elem().Kind() {

	case reflect.Array, reflect.Slice:

		// Iterate through the Array/Slice and randomize each of the element
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
				v.SetMapIndex(e, reflect.ValueOf(getRandomInt()))

			case reflect.String:
				v.SetMapIndex(e, reflect.ValueOf(getRandomString(len(v.String()))))

			case reflect.Struct:
				randomValueGenerator(v)

			case reflect.Float64:
				v.SetMapIndex(e, reflect.ValueOf(getRandomFloat()))
			}
		}
	case reflect.Struct:

		// run it through randomValueGenerator
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
	for i := 0; i < mutable.NumField(); i++ {

		switch mutable.Field(i).Kind() {

		case reflect.Struct:
			randomValueGenerator(mutable.Field(i))

		case reflect.Int:
			mutable.Field(i).SetInt(getRandomInt())

		case reflect.Float64:
			mutable.Field(i).SetFloat(getRandomFloat())

		case reflect.String:
			mutable.Field(i).SetString(getRandomString(len(mutable.Field(i).String())))

		case reflect.Map:
			for _, e := range mutable.Field(i).MapKeys() {
				v := mutable.Field(i).MapIndex(e)
				switch v.Elem().Kind() {
				case reflect.Int:
					mutable.Field(i).SetMapIndex(e, reflect.ValueOf(getRandomInt()))

				case reflect.String:
					mutable.Field(i).SetMapIndex(e, reflect.ValueOf(getRandomString(len(v.String()))))

				case reflect.Struct:
					randomValueGenerator(v)

				case reflect.Float64:
					mutable.Field(i).SetMapIndex(e, reflect.ValueOf(getRandomFloat()))
				}
			}

		case reflect.Array, reflect.Slice:

			for j := 0; j < mutable.Field(i).Len(); j++ {
				randomValueGenerator(mutable.Field(i).Index(j))
			}
		}
	}
}

// getRandomInt ...
// returns a random int64
func getRandomInt() int64 {

	return int64(rand.Int())
}

// getRandomFloat ...
// returns a random float64
func getRandomFloat() float64 {

	return rand.Float64()
}

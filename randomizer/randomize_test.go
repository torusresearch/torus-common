package randomizer

import (
	"reflect"
	"strconv"
	"testing"
)

type Simple struct {
	Integer int
	String  string
	Float   float64
}

type Complex struct {
	Struct  Simple
	Map     map[string]interface{}
	Boolean bool
	Array   []Simple
}

func TestRandomize(t *testing.T) {

	var (
		originalStructure Complex
	)

	testMap := make(map[string]interface{})
	testMap["map1"] = 100
	testMap["map2"] = 23.78
	testMap["map3"] = "this is a random string"

	testStructure := Complex{
		Struct: Simple{
			Integer: 200,
			String:  "this is also a random string",
			Float:   23.220,
		},
		Map:     testMap,
		Boolean: false,
		Array: []Simple{
			{Integer: 200, String: "a random string", Float: 23.33},
			{Integer: 253, String: "another random string", Float: 63.33},
			{Integer: 243, String: "another another random string", Float: 993.33},
		},
	}

	originalStructure = copyTestStruct(testStructure)

	err := Randomize(&testStructure)

	if err != nil {
		t.Log("shouldn't throw any errors")
		t.Error("expected nil, got " + err.Error())
		t.Fail()
		return
	}

	// Tests for Complex.Boolean
	if originalStructure.Boolean != testStructure.Boolean {

		// The bool shouldn't be part of the Randomize
		t.Log("shouldn't randomize bool")
		t.Error("Expected " + strconv.FormatBool(originalStructure.Boolean) + " got " + strconv.FormatBool(testStructure.Boolean))
		t.Fail()
	}

	// MARK: Tests for Complex.Map
	for key, value := range originalStructure.Map {

		randomizedValue, ok := testStructure.Map[key]

		if !ok {
			// Key doesn't exist error
			t.Log("the randomized map should have values for all the specified keys")
			t.Error("couldn't find the value for key: " + key)
			t.Fail()
		}

		if reflect.TypeOf(randomizedValue).Kind() != reflect.TypeOf(value).Kind() {

			// The type of the value is changed after Randomize
			if value == randomizedValue {
				t.Log("the randomized value should be different than the original value")
				t.Error("expected random value. Got " + randomizedValue.(string))
				t.Fail()
			}
		}

		// Tests for Values in a map

		switch reflect.TypeOf(randomizedValue).Kind() {

		case reflect.Struct:
			// Tests for the map value when it is a Struct
			originalMapStruct := value.(Simple)
			randomizedMapStruct := value.(Simple)

			if originalMapStruct.String == randomizedMapStruct.String {

				t.Log("the randomized value should be different than the original value")
				t.Error("expected random value. Got " + randomizedMapStruct.String)
				t.Fail()
			}

			if originalMapStruct.Float == randomizedMapStruct.Float {

				t.Log("the randomized value should be different than the original value")
				t.Error("expected random value. Got " + strconv.FormatFloat(randomizedMapStruct.Float, 'f', 'g', 64))
				t.Fail()
			}

			if originalMapStruct.Integer == randomizedMapStruct.Integer {

				t.Log("the randomized value should be different than the original value")
				t.Error("expected random value. Got " + strconv.FormatInt(int64(randomizedMapStruct.Integer), 10))
				t.Fail()
			}

		case reflect.Int:
			if randomizedValue.(int) != value.(int) {
				t.Log("the randomized value should be different than the original value")
				t.Error("expected random value. Got " + strconv.FormatInt(randomizedValue.(int64), 10))
				t.Fail()
			}

		case reflect.String:
			if randomizedValue.(string) != value.(string) {
				t.Log("the randomized value should be different than the original value")
				t.Error("expected random value. Got " + randomizedValue.(string))
				t.Fail()
			}

		case reflect.Float64:
			if randomizedValue.(float64) != value.(float64) {
				t.Log("the randomized value should be different than the original value")
				t.Error("expected random value. Got " + strconv.FormatFloat(randomizedValue.(float64), 'f', 'g', 64))
				t.Fail()
			}

		case reflect.Bool:
			if randomizedValue.(bool) != value.(bool) {
				t.Log("shouldn't randomize bool")
				t.Error("Expected " + strconv.FormatBool(value.(bool)) + " got " + strconv.FormatBool(randomizedValue.(bool)))
				t.Fail()
			}
		}
	}

	// MARK: Tests for Complex.Struct

	if originalStructure.Struct.String == testStructure.Struct.String {

		t.Log("the randomized value should be different than the original value")
		t.Error("expected random value. Got " + testStructure.Struct.String)
		t.Fail()
	}

	if originalStructure.Struct.Integer == testStructure.Struct.Integer {

		t.Log("the randomized value should be different than the original value")
		t.Error("expected random value. Got " + strconv.FormatInt(int64(testStructure.Struct.Integer), 10))
		t.Fail()
	}

	if originalStructure.Struct.Float == testStructure.Struct.Float {

		t.Log("the randomized value should be different than the original value")
		t.Error("expected random value. Got " + strconv.FormatFloat(testStructure.Struct.Float, 'f', 'g', 64))
		t.Fail()
	}

	// MARK: Tests for Complex.Array
	if len(originalStructure.Array) != len(testStructure.Array) {
		t.Log("the length of the randomized array should be the same as the original array")
		t.Error("expected length : " + strconv.Itoa(len(originalStructure.Array)) + " Got " + strconv.Itoa(len(testStructure.Array)))
		t.Fail()
	}

	for i := 0; i < len(originalStructure.Array); i++ {

		if originalStructure.Array[i].Float == testStructure.Array[i].Float {
			t.Log("the randomized value should be different than the original value")
			t.Error("expected random value. Got " + strconv.FormatFloat(testStructure.Array[i].Float, 'f', 'g', 64))
			t.Fail()
		}

		if originalStructure.Array[i].Integer == testStructure.Array[i].Integer {
			t.Log("the randomized value should be different than the original value")
			t.Error("expected random value. Got " + strconv.Itoa(testStructure.Array[i].Integer))
			t.Fail()
		}

		if originalStructure.Array[i].String == testStructure.Array[i].String {
			t.Log("the randomized value should be different than the original value")
			t.Error("expected random value. Got " + testStructure.Array[i].String)
			t.Fail()
		}
	}
}

// MARK: Helper methods
// copyTestStruct ...
// Creates a copy of the struct
// Didn't wanna use reflect while testing reflect
func copyTestStruct(original Complex) Complex {

	duplicate := Complex{
		Struct:  original.Struct,
		Map:     original.Map,
		Boolean: original.Boolean,
	}

	duplicate.Array = make([]Simple, len(original.Array))

	copy(duplicate.Array, original.Array)

	return duplicate
}

// MARK: Benchmark

func BenchmarkRandomize(b *testing.B) {
	testMap := make(map[string]interface{})
	testMap["map1"] = 100
	testMap["map2"] = 23.78
	testMap["map3"] = "this is a random string"

	testStructure := Complex{
		Struct: Simple{
			Integer: 200,
			String:  "this is also a random string",
			Float:   23.220,
		},
		Map:     testMap,
		Boolean: false,
		Array: []Simple{
			{Integer: 200, String: "a random string", Float: 23.33},
			{Integer: 253, String: "another random string", Float: 63.33},
			{Integer: 243, String: "another another random string", Float: 993.33},
		},
	}

	for i := 0; i < 100; i++ {
		Randomize(&testStructure)
	}
}

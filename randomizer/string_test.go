package randomizer

import "testing"

//
const n = 16

// BenchmarkGetRandomString ...
func BenchmarkGetRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getRandomString(n)
	}
}

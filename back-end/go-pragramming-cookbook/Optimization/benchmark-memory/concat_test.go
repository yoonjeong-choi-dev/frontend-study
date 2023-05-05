package benchmark_memory

import "testing"

func BenchmarkConcat(b *testing.B) {
	b.Run("one", func(b *testing.B) {
		one := []string{"1"}
		for i := 0; i < b.N; i++ {
			concat(one...)
		}
	})

	b.Run("five", func(b *testing.B) {
		five := []string{"1", "2", "3", "4", "5"}
		for i := 0; i < b.N; i++ {
			concat(five...)
		}
	})

	b.Run("ten", func(b *testing.B) {
		ten := []string{"1", "2", "3", "4", "5",
			"6", "7", "8", "9", "10"}
		for i := 0; i < b.N; i++ {
			concat(ten...)
		}
	})
}

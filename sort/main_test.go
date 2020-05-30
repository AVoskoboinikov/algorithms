package sort

import "testing"

func BenchmarkSort(b *testing.B) {
	arr := readArr()

	b.Run("basic", func(b *testing.B) {
		b.ResetTimer()

		for n := 0; n < b.N; n++ {
			_ = basicSort(arr)
		}
	})

	b.Run("merge", func(b *testing.B) {
		b.ResetTimer()

		for n := 0; n < b.N; n++ {
			_, _ = mergeSort(arr)
		}
	})
}

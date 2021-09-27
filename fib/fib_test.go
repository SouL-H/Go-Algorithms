package main

import "testing"

//  214250              5630 ns/op               0 B/op          0 allocs/op
func BenchmarkFib(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = Fib(15)
	}
}
// 517987660                2.061 ns/op           0 B/op          0 allocs/op
func BenchmarkFibOptimize(b *testing.B) {
	n := 15
	memo := make([]int, n)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = FibOptimzie(n, memo)
	}
}

// 39473767                30.05 ns/op            0 B/op          0 allocs/op
func BenchmarkFibBottomUp(b *testing.B){
	n:=15
	k:= make([]int,n+1)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = FibBottomUp(n,k)
		
	}
}
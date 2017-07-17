package main

import "testing"

var a = ArrayInit()

func BenchmarkArrayInit(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ArrayInit()
	}
}

func BenchmarkArrayModify(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := range a {
			a[i].Value += i
		}
	}
}

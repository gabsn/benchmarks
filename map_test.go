package main

import "testing"

var m = MapInitClassic()

func BenchmarkMapInitClassic(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MapInitClassic()
	}
}

func BenchmarkMapInitIterate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MapInitIterate()
	}
}

func BenchmarkMapModify(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := range m {
			m[i] += 1
		}
	}
}

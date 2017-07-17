package main

import (
	"reflect"
	"testing"
)

var s = StructInitClassic()
var s2 = StructInitClassic()

func BenchmarkStructInitClassic(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StructInitClassic()
	}
}

func BenchmarkStructInitReflect(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StructInitReflect()
	}
}

func BenchmarkStructInitIterate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StructInitIterate()
	}
}

func BenchmarkStructModify(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s2.A += 1
		s2.B += 1
		s2.C += 1
		s2.D += 1
		s2.E += 1
		s2.F += 1
	}
}

func BenchmarkStructModifyReflect(b *testing.B) {
	for n := 0; n < b.N; n++ {
		v := reflect.ValueOf(s).Elem()
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			field.SetInt(field.Int() + 1)
		}
	}
}

func BenchmarkStructTags(b *testing.B) {
	for n := 0; n < b.N; n++ {
		t := reflect.TypeOf(s).Elem()
		for i := 0; i < t.NumField(); i++ {
			_ = string(t.Field(i).Tag)
		}
	}
}

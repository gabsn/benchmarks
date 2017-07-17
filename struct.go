package main

import "reflect"

type S struct {
	A int64
	B int64
	C int64
	D int64
	E int64
	F int64
}

func StructInitClassic() *S {
	// Allocate in memory
	s := &S{}

	// Initialize
	s.A = 0
	s.B = 1
	s.C = 2
	s.D = 3
	s.E = 4
	s.F = 5

	return s
}

func StructInitReflect() *S {
	// Allocate in memory
	s := &S{}
	v := reflect.ValueOf(s).Elem()

	// Initialize with reflect
	v.FieldByName("A").SetInt(0)
	v.FieldByName("B").SetInt(1)
	v.FieldByName("C").SetInt(2)
	v.FieldByName("D").SetInt(3)
	v.FieldByName("E").SetInt(4)
	v.FieldByName("F").SetInt(5)

	return s
}

func StructInitIterate() *S {
	// Allocate in memory
	s := &S{}
	v := reflect.ValueOf(s).Elem()

	// Iterate over the values
	for i := 0; i < v.NumField(); i++ {
		v.Field(i).SetInt(int64(i))
	}

	return s
}

package main

type A []*F

type F struct {
	Name  string
	Value int
}

func newF(name string, value int) *F {
	return &F{name, value}
}

func ArrayInit() A {
	// Allocate in memory
	a := make(A, 6)

	// Iterate to initalize
	for i := range a {
		a[i] = newF(alpha[i], i)
	}

	return a
}

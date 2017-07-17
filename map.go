package main

func MapInitClassic() map[string]int {
	// Allocate in memory
	m := make(map[string]int, 6)

	// Initialize
	m["A"] = 0
	m["B"] = 1
	m["C"] = 2
	m["D"] = 3
	m["E"] = 4
	m["F"] = 5

	return m
}

func MapInitIterate() map[string]int {
	// Allocate in memory
	m := make(map[string]int, 6)

	// Iterate over the values
	for i := 0; i < 6; i++ {
		m[alpha[i]] = i
	}

	return m
}

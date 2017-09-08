# benchmarks
Go benchmarks to make smart choices.

## Run benchmarks

`go test -bench "Struct|Map|Array" -benchmem`

- Example of output:
```
BenchmarkArrayInit-4             	 5000000	       291 ns/op	     240 B/op	       7 allocs/op
BenchmarkArrayModify-4           	300000000	         5.86 ns/op	       0 B/op	       0 allocs/op
BenchmarkMapInitClassic-4        	 5000000	       289 ns/op	       0 B/op	       0 allocs/op
BenchmarkMapModify-4             	 3000000	       410 ns/op	       0 B/op	       0 allocs/op
BenchmarkStructInitClassic-4     	300000000	         4.03 ns/op	       0 B/op	       0 allocs/op
BenchmarkStructModify-4          	1000000000	         2.95 ns/op	       0 B/op	       0 allocs/op
```

## Conclusion

In this case, we see that struct outperforms the map and the array of structs.
However the time to modify array of structs is close to the one to modify the struct itself.
So we should consider using an array of structs instead of a struct to avoid duplicating code.

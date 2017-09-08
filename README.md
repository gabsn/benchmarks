# benchmarks
Some go benchmarks I used to make smart implementation choices.

## sync.Pool VS channel
### Run
```
go test -bench=Pool -benchmem
```
### Output
```
BenchmarkOnePool-4                 	 3000000	       436 ns/op	       0 B/op	       0 allocs/op
BenchmarkOneChanPool-4             	 2000000	       844 ns/op	       0 B/op	       0 allocs/op
BenchmarkTwoPoolNoCopy-4           	 2000000	       647 ns/op	       0 B/op	       0 allocs/op
BenchmarkTwoChanPoolNoCopy-4       	 1000000	      1647 ns/op	       0 B/op	       0 allocs/op
BenchmarkTwoPoolWriteTo-4          	 2000000	       773 ns/op	       0 B/op	       0 allocs/op
BenchmarkTwoChanPoolWriteTo-4      	 1000000	      1736 ns/op	       0 B/op	       0 allocs/op
BenchmarkTwoPoolCopy-4             	 2000000	       940 ns/op	       0 B/op	       0 allocs/op
BenchmarkTwoChanPoolCopy-4         	 1000000	      1899 ns/op	       0 B/op	       0 allocs/op
BenchmarkTwoPoolCopyBuffer-4       	 2000000	       985 ns/op	       0 B/op	       0 allocs/op
BenchmarkTwoChanPoolCopyBuffer-4   	 1000000	      1924 ns/op	       0 B/op	       0 allocs/op
```
### Conclusion
We clearly see that we should use `sync.Pool` instead of `chan *bytes.Buffer` to implement the pool in this case. Then, we see that we should use `WriteTo` method instead of `Copy` or `CopyBuffer` when we use two pools. Finally, we see that we are _twice slower when using two pools instead of one pool_. We should also note that in both implementations, we manage to avoid allocations when dealing with those buffers.

## Struct VS Array VS Map
### Run
`go test -bench "Struct|Map|Array" -benchmem`
### Output
```
BenchmarkArrayInit-4             	 5000000	       291 ns/op	     240 B/op	       7 allocs/op
BenchmarkArrayModify-4           	300000000	         5.86 ns/op	       0 B/op	       0 allocs/op
BenchmarkMapInitClassic-4        	 5000000	       289 ns/op	       0 B/op	       0 allocs/op
BenchmarkMapModify-4             	 3000000	       410 ns/op	       0 B/op	       0 allocs/op
BenchmarkStructInitClassic-4     	300000000	         4.03 ns/op	       0 B/op	       0 allocs/op
BenchmarkStructModify-4          	1000000000	         2.95 ns/op	       0 B/op	       0 allocs/op
```
### Conclusion
In this case, we see that struct outperforms the map and the array of structs.
However the time to modify array of structs is close to the one to modify the struct itself.
So we should consider using an array of structs instead of a struct to avoid duplicating code.

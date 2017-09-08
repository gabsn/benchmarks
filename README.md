# benchmarks
Some go benchmarks I used to make smart implementation choices.

## sync.Pool VS channel
### Run
```
go test -bench=Pool -benchmem
```
### Output 
- **Synchronous Operations**
```
BenchmarkOnePool-4                           	30000000	        41.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkOneChanPool-4                       	20000000	        84.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkTwoPoolNoCopy-4                     	20000000	        62.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkTwoChanPoolNoCopy-4                 	10000000	       171 ns/op	       0 B/op	       0 allocs/op
BenchmarkTwoPoolWriteTo-4                    	20000000	        77.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkTwoChanPoolWriteTo-4                	10000000	       171 ns/op	       0 B/op	       0 allocs/op
BenchmarkTwoPoolCopy-4                       	20000000	        94.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkTwoChanPoolCopy-4                   	10000000	       189 ns/op	       0 B/op	       0 allocs/op
BenchmarkTwoPoolCopyBuffer-4                 	20000000	        91.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkTwoChanPoolCopyBuffer-4             	10000000	       201 ns/op	       0 B/op	       0 allocs/op
```

- **Concurrent Operations**
```
BenchmarkConcurrentOnePool-4                 	  300000	      3965 ns/op	       0 B/op	       0 allocs/op
BenchmarkConcurrentOneChanPool-4             	  300000	      5129 ns/op	       0 B/op	       0 allocs/op
BenchmarkConcurrentTwoPoolNoCopy-4           	  300000	      4605 ns/op	       0 B/op	       0 allocs/op
BenchmarkConcurrentTwoChanPoolNoCopy-4       	  200000	      6682 ns/op	       0 B/op	       0 allocs/op
BenchmarkConcurrentTwoPoolWriteTo-4          	  300000	      4835 ns/op	       0 B/op	       0 allocs/op
BenchmarkConcurrentTwoChanPoolWriteTo-4      	  200000	      6903 ns/op	       0 B/op	       0 allocs/op
BenchmarkConcurrentTwoPoolCopy-4             	  300000	      4953 ns/op	       0 B/op	       0 allocs/op
BenchmarkConcurrentTwoChanPoolCopy-4         	  200000	      7325 ns/op	       0 B/op	       0 allocs/op
BenchmarkConcurrentTwoPoolCopyBuffer-4       	  300000	      5691 ns/op	       0 B/op	       0 allocs/op
BenchmarkConcurrentTwoChanPoolCopyBuffer-4   	  200000	      7550 ns/op	       0 B/op	       0 allocs/op
```
### Conclusion
- **Synchronous benchmark**
We clearly see that we should use `sync.Pool` instead of `chan *bytes.Buffer` to implement the pool in this case. Then, we see that we should use `WriteTo` method instead of `Copy` or `CopyBuffer` when we use two pools. Finally, we see that we are _twice slower when using two pools instead of one pool_. We should also note that in both implementations, we manage to avoid allocations when dealing with those buffers.

- **Concurrent benchmark**
In this benchmark, I spawned 10 concurrent goroutines for each operation trying to access the pool, write into a buffer and put it back to the pool. In this case, we see that the channel implementation of the buffer pool is _only 20% slower_ than the `sync.Pool` implementation. Moreover copying the buffers two the second pool with the `WriteTo` method only takes _20% more time than when we do not use the second pool_, which is more acceptable than the 100% with synchronous operations. 

You can see below the overhead of using the `WriteTo` method to copy a buffer from the first pool to a buffer from the second pool compared with only using one pool.

![WriteTo overhead]()

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

package main

import (
	"io"
	"sync"
	"testing"
)

const MaxGoroutines = 25

func BenchmarkOnePool(b *testing.B) {
	pool := NewBufferPool()
	for n := 0; n < b.N; n++ {
		b := pool.Get()
		b.Reset()
		b.Write([]byte("Hello World!"))
		pool.Put(b)
	}
}

func BenchmarkOneChanPool(b *testing.B) {
	pool := NewBufferChanPool()
	for n := 0; n < b.N; n++ {
		b := pool.Get()
		b.Reset()
		b.Write([]byte("Hello World!"))
		pool.Put(b)
	}
}

func BenchmarkTwoPoolNoCopy(b *testing.B) {
	pool1 := NewBufferPool()
	pool2 := NewBufferPool()
	for n := 0; n < b.N; n++ {
		b1 := pool1.Get()
		b2 := pool2.Get()
		b1.Reset()
		b2.Reset()
		b1.Write([]byte("Hello World!"))
		pool1.Put(b1)
		pool2.Put(b2)
	}
}

func BenchmarkTwoChanPoolNoCopy(b *testing.B) {
	pool1 := NewBufferChanPool()
	pool2 := NewBufferChanPool()
	for n := 0; n < b.N; n++ {
		b1 := pool1.Get()
		b2 := pool2.Get()
		b1.Reset()
		b2.Reset()
		b1.Write([]byte("Hello World!"))
		pool1.Put(b1)
		pool2.Put(b2)
	}
}

func BenchmarkTwoPoolWriteTo(b *testing.B) {
	pool1 := NewBufferPool()
	pool2 := NewBufferPool()
	for n := 0; n < b.N; n++ {
		b1 := pool1.Get()
		b2 := pool2.Get()
		b1.Reset()
		b2.Reset()
		b1.Write([]byte("Hello World!"))
		b1.WriteTo(b2)
		pool1.Put(b1)
		pool2.Put(b2)
	}
}

func BenchmarkTwoChanPoolWriteTo(b *testing.B) {
	pool1 := NewBufferChanPool()
	pool2 := NewBufferChanPool()
	for n := 0; n < b.N; n++ {
		b1 := pool1.Get()
		b2 := pool2.Get()
		b1.Reset()
		b2.Reset()
		b1.Write([]byte("Hello World!"))
		b1.WriteTo(b2)
		pool1.Put(b1)
		pool2.Put(b2)
	}
}

func BenchmarkTwoPoolCopy(b *testing.B) {
	pool1 := NewBufferPool()
	pool2 := NewBufferPool()
	for n := 0; n < b.N; n++ {
		b1 := pool1.Get()
		b2 := pool2.Get()
		b1.Reset()
		b2.Reset()
		b1.Write([]byte("Hello World!"))
		io.Copy(b2, b1)
		pool1.Put(b1)
		pool2.Put(b2)
	}
}

func BenchmarkTwoChanPoolCopy(b *testing.B) {
	pool1 := NewBufferChanPool()
	pool2 := NewBufferChanPool()
	for n := 0; n < b.N; n++ {
		b1 := pool1.Get()
		b2 := pool2.Get()
		b1.Reset()
		b2.Reset()
		b1.Write([]byte("Hello World!"))
		io.Copy(b2, b1)
		pool1.Put(b1)
		pool2.Put(b2)
	}
}

func BenchmarkTwoPoolCopyBuffer(b *testing.B) {
	pool1 := NewBufferPool()
	pool2 := NewBufferPool()
	buf := make([]byte, 8)
	for n := 0; n < b.N; n++ {
		b1 := pool1.Get()
		b2 := pool2.Get()
		b1.Reset()
		b2.Reset()
		b1.Write([]byte("Hello World!"))
		io.CopyBuffer(b2, b1, buf)
		pool1.Put(b1)
		pool2.Put(b2)
	}
}

func BenchmarkTwoChanPoolCopyBuffer(b *testing.B) {
	pool1 := NewBufferChanPool()
	pool2 := NewBufferChanPool()
	buf := make([]byte, 8)
	for n := 0; n < b.N; n++ {
		b1 := pool1.Get()
		b2 := pool2.Get()
		b1.Reset()
		b2.Reset()
		b1.Write([]byte("Hello World!"))
		io.CopyBuffer(b2, b1, buf)
		pool1.Put(b1)
		pool2.Put(b2)
	}
}

func BenchmarkConcurrentOnePool(b *testing.B) {
	var wg sync.WaitGroup
	pool := NewBufferPool()
	for n := 0; n < b.N; n++ {
		for i := 0; i < MaxGoroutines; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				b := pool.Get()
				b.Reset()
				b.Write([]byte("Hello World!"))
				pool.Put(b)
			}()
		}
		wg.Wait()
	}
}

func BenchmarkConcurrentOneChanPool(b *testing.B) {
	var wg sync.WaitGroup
	pool := NewBufferChanPool()
	for n := 0; n < b.N; n++ {
		for i := 0; i < MaxGoroutines; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				b := pool.Get()
				b.Reset()
				b.Write([]byte("Hello World!"))
				pool.Put(b)
			}()
		}
		wg.Wait()
	}
}

func BenchmarkConcurrentTwoPoolNoCopy(b *testing.B) {
	var wg sync.WaitGroup
	pool1 := NewBufferPool()
	pool2 := NewBufferPool()
	for n := 0; n < b.N; n++ {
		for i := 0; i < MaxGoroutines; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				b1 := pool1.Get()
				b2 := pool2.Get()
				b1.Reset()
				b2.Reset()
				b1.Write([]byte("Hello World!"))
				pool1.Put(b1)
				pool2.Put(b2)
			}()
		}
		wg.Wait()
	}
}

func BenchmarkConcurrentTwoChanPoolNoCopy(b *testing.B) {
	var wg sync.WaitGroup
	pool1 := NewBufferChanPool()
	pool2 := NewBufferChanPool()
	for n := 0; n < b.N; n++ {
		for i := 0; i < MaxGoroutines; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				b1 := pool1.Get()
				b2 := pool2.Get()
				b1.Reset()
				b2.Reset()
				b1.Write([]byte("Hello World!"))
				pool1.Put(b1)
				pool2.Put(b2)
			}()
		}
		wg.Wait()
	}
}

func BenchmarkConcurrentTwoPoolWriteTo(b *testing.B) {
	var wg sync.WaitGroup
	pool1 := NewBufferPool()
	pool2 := NewBufferPool()
	for n := 0; n < b.N; n++ {
		for i := 0; i < MaxGoroutines; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				b1 := pool1.Get()
				b2 := pool2.Get()
				b1.Reset()
				b2.Reset()
				b1.Write([]byte("Hello World!"))
				b1.WriteTo(b2)
				pool1.Put(b1)
				pool2.Put(b2)
			}()
		}
		wg.Wait()
	}
}

func BenchmarkConcurrentTwoChanPoolWriteTo(b *testing.B) {
	var wg sync.WaitGroup
	pool1 := NewBufferChanPool()
	pool2 := NewBufferChanPool()
	for n := 0; n < b.N; n++ {
		for i := 0; i < MaxGoroutines; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				b1 := pool1.Get()
				b2 := pool2.Get()
				b1.Reset()
				b2.Reset()
				b1.Write([]byte("Hello World!"))
				b1.WriteTo(b2)
				pool1.Put(b1)
				pool2.Put(b2)
			}()
		}
		wg.Wait()
	}
}

func BenchmarkConcurrentTwoPoolCopy(b *testing.B) {
	var wg sync.WaitGroup
	pool1 := NewBufferPool()
	pool2 := NewBufferPool()
	for n := 0; n < b.N; n++ {
		for i := 0; i < MaxGoroutines; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				b1 := pool1.Get()
				b2 := pool2.Get()
				b1.Reset()
				b2.Reset()
				b1.Write([]byte("Hello World!"))
				io.Copy(b2, b1)
				pool1.Put(b1)
				pool2.Put(b2)
			}()
		}
		wg.Wait()
	}
}

func BenchmarkConcurrentTwoChanPoolCopy(b *testing.B) {
	var wg sync.WaitGroup
	pool1 := NewBufferChanPool()
	pool2 := NewBufferChanPool()
	for n := 0; n < b.N; n++ {
		for i := 0; i < MaxGoroutines; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				b1 := pool1.Get()
				b2 := pool2.Get()
				b1.Reset()
				b2.Reset()
				b1.Write([]byte("Hello World!"))
				io.Copy(b2, b1)
				pool1.Put(b1)
				pool2.Put(b2)
			}()
		}
		wg.Wait()
	}
}

func BenchmarkConcurrentTwoPoolCopyBuffer(b *testing.B) {
	var wg sync.WaitGroup
	pool1 := NewBufferPool()
	pool2 := NewBufferPool()
	buf := make([]byte, 8)
	for n := 0; n < b.N; n++ {
		for i := 0; i < MaxGoroutines; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				b1 := pool1.Get()
				b2 := pool2.Get()
				b1.Reset()
				b2.Reset()
				b1.Write([]byte("Hello World!"))
				io.CopyBuffer(b2, b1, buf)
				pool1.Put(b1)
				pool2.Put(b2)
			}()
		}
		wg.Wait()
	}
}

func BenchmarkConcurrentTwoChanPoolCopyBuffer(b *testing.B) {
	var wg sync.WaitGroup
	pool1 := NewBufferChanPool()
	pool2 := NewBufferChanPool()
	buf := make([]byte, 8)
	for n := 0; n < b.N; n++ {
		for i := 0; i < MaxGoroutines; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				b1 := pool1.Get()
				b2 := pool2.Get()
				b1.Reset()
				b2.Reset()
				b1.Write([]byte("Hello World!"))
				io.CopyBuffer(b2, b1, buf)
				pool1.Put(b1)
				pool2.Put(b2)
			}()
		}
		wg.Wait()
	}
}

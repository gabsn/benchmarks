package main

import (
	"bytes"
	"sync"
)

type BufferPool struct {
	*sync.Pool
}

func NewBufferPool() *BufferPool {
	pool := &sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
	return &BufferPool{pool}
}

func (p *BufferPool) Get() *bytes.Buffer {
	return p.Pool.Get().(*bytes.Buffer)
}

type BufferChanPool struct {
	pool chan *bytes.Buffer
}

func NewBufferChanPool() *BufferChanPool {
	pool := make(chan *bytes.Buffer, 100)
	return &BufferChanPool{pool}
}

func (c *BufferChanPool) Get() *bytes.Buffer {
	var b *bytes.Buffer
	select {
	case b = <-c.pool:
		// Got one; nothing more to do.
	default:
		// None free, so allocate a new one.
		b = new(bytes.Buffer)
	}
	return b
}

func (p *BufferChanPool) Put(b *bytes.Buffer) {
	p.pool <- b
}

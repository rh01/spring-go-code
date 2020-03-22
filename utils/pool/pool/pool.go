package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

// Pool 管理一组可以安全地在多个 goroutine 中共享的资源
// 被管理的资源必须实现了io.Closer接口
type Pool struct {
	m         sync.Mutex                // 互斥锁，用来goroutine对资源池的互斥访问
	resources chan io.Closer            // 资源池中的资源，必须实现io.Closer接口
	factory   func() (io.Closer, error) // 分配资源的函数
	closed    bool                      // 判断当前的资源池是否关闭
}

// ErrPoolClosed 表示请求了一个已经关闭的资源池
var ErrPoolClosed = errors.New("pool closed")

// New 创建一个管理资源的池子，这个池需要一个分配新资源 的函数，
// 并规定这个池子的大小
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("Size value too small")
	}

	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

// Acquire 从资源池子里面取出一个资源
func (p *Pool) Acquire() (io.Closer, error) {
	// 基本的逻辑如下：若资源池已经关闭，将会返回一个错误，否则阻塞
	select {
	// 检查当前是否有空闲的资源
	case r, ok := <-p.resources:
		log.Println("Acquired: ", "Shared resources.")
		// 若资源池已经关闭，则会返回一个错误
		if !ok {
			return nil, ErrPoolClosed
		}

		return r, nil

	// 因为没有空闲的资源，这时候需要提供一个新的资源，
	// 这里的意思是生产一个新的资源的意思吗？
	default:
		log.Println("Acquired：", "New Resource")
		return p.factory()
	}
}

// Release 释放一个资源到资源池中
func (p *Pool) Release(r io.Closer) {
	// 保证互斥操作
	p.m.Lock()
	defer p.m.Unlock()

	// 如果当前的资源池已经关闭，需要释放的资源也要关闭，即
	// 销毁，然后返回即可
	if p.closed {
		r.Close()
		return
	}

	// 然后阻塞在这里
	select {
	// 试图放入资源池，因为这时候可能会通道满了
	case p.resources <- r:
		// 放入成功
		log.Println("Resource：", "In Queue")

	// 如果队列已经满了，则关闭这个资源
	default:
		log.Println("Release: ", "Closing")
		r.Close()
	}
}

// Close  会让资源池停止工作，并关闭所有现有的资源
func (p *Pool) Close() {
	// 保证互斥
	p.m.Lock()
	defer p.m.Unlock()

	// 判断是否关闭
	if p.closed {
		return
	}

	// 蒋池子关闭，这时候释放资源的将会自动释放
	p.closed = true

	// 在清空通道里面的资源之前，江通道关闭
	// 如果不这样做，会发生死锁
	close(p.resources)

	// 关闭资源
	for r := range p.resources {
		r.Close()
	}
}

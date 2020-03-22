package main

import (
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"github.com/rh01/leetcode/utils/pool/pool"
)

const (
	maxGoroutines   = 25 // 要使用的goroutine数量
	pooledResources = 2  // 池中资源的数量
)

// dbConnection 是用来模拟的资源
type dbConnection struct {
	ID int32
}

// dbCOnnection 要实现Close的方法,即实现了
// io.CLoser接口，以方便Pool的管理
func (db *dbConnection) Close() error {
	log.Println("Close: connection: ", db.ID)
	return nil
}

// idCounter 用来给每个链接分配一个独一无二的id
var idCounter int32

// createConnection 是一个工厂函数，
// 当需要一个新的连接时，资源池会调用这个函数
func createConnection() (io.Closer, error) {
	// 保证互斥
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New connection ", id)

	return &dbConnection{id}, nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	// 创建用来管理链接的池
	p, err := pool.New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}

	// 使用池子中链接进行查询
	for i := 0; i < maxGoroutines; i++ {
		// 每个goroutine需要复制一份要查询的副本，不然所有的查询会共享
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(i)
	}

	// 等待结果
	wg.Wait()

	// 关闭资源池
	log.Println("shutdown Program.")
	p.Close()
}

// performQueries 用来测试链接的资源池
func performQueries(q int, p *pool.Pool) {
	// 从池子里面获取一个链接
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}

	// 将该链接放回池子中
	defer p.Release(conn)

	// 用等待来模拟查询响应
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("QID[%d] CID[%d]\n", q, conn.(*dbConnection).ID)
}

package main

import (
	"log"
	"os"
	"time"

	"github.com/rh01/leetcode/utils/runner/runner"
)

// timeout 规定了多长时间内程序结束
const timeout = 5 * time.Second

func main() {
	log.Println("Starting work.")

	// 为本次执行分配超时时间
	r := runner.New(timeout)

	// 加入要执行的任务
	r.Add(creatTask(), creatTask(), creatTask())

	// 执行任务并处理结果
	if err := r.Start(); r != nil {
		switch err {
		case runner.ErrInterrupt:
			log.Println("Interrupt Error")
			os.Exit(2)

		case runner.ErrTimeout:
			log.Println("Timeout")
			os.Exit(1)
		}
	}

	log.Println("Process end.")
}

//  这是一个闭包函数
func creatTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d\n", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}

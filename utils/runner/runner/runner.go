package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// Runner 在给定的超时时间里执行一组任务
type Runner struct {
	// interrupt 通道报告从操作系统接收的信号
	interrupt chan os.Signal

	// completed 通道用来报告任务是否完成
	completed chan error

	// timeout 通道用来报告处理任务已经超时
	timeout <-chan time.Time

	// tasks 函数表示一组要顺序执行的任务函数
	tasks []func(int)
}

// ErrTimeout 会在任务执行超时时返回
var ErrTimeout = errors.New("receive timeout")

// ErrInterrupt 会在操作系统发送信号时返回
var ErrInterrupt = errors.New("recerive interrupt")

// New 新建一个准备使用的Runner对象
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		completed: make(chan error),
		timeout:   time.After(d),
	}
}

// Add 将任务附加到Runner上
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// Start 开始执行任务
func (r *Runner) Start() error {
	// 希望接收所有终端信号
	signal.Notify(r.interrupt, os.Interrupt)

	// 用不同的go routine执行不同的任务
	go func() {
		r.completed <- r.run()
	}()

	select {
	case err := <-r.completed:
		return err

	case <-r.timeout:
		return ErrInterrupt
	}
}

func (r *Runner) run() error {

	for id, task := range r.tasks {
		// 检测操作系统的终端信号
		if r.gotInterrupt() {
			return ErrInterrupt
		}

		// 执行已经注册的任务
		task(id)
	}

	return nil
}

// 检查操作系统的中断信号
func (r *Runner) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		// 停止接受后续的信号
		signal.Stop(r.interrupt)
		return true

	// 默认继续执行
	default:
		return false
	}
}

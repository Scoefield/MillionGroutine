package main

import (
	"fmt"
	"practicProject/myTest/milliongroutine/handler"
	"runtime"
	"time"
)

type Score struct {
	Num int
}

func (s *Score) Do() {
	// do处理的逻辑
	fmt.Println("num:", s.Num)
	time.Sleep(1 * 1 * time.Second)
}

func main() {
	num := 100 * 100 * 20
	// debug.SetMaxThreads(num + 1000) //设置最大线程数
	// 注册工作池，传入任务
	// 参数1 worker并发个数
	p := handler.NewWorkerPool(num)
	p.Run()

	// 这里根据实际情况，改为自己的需要处理的逻辑（做相应的修改）
	dataNum := 100 * 100 * 100 * 100
	go func() {
		for i := 1; i <= dataNum; i++ {
			sc := &Score{Num: i}
			p.JobQueue <- sc
			time.Sleep(500* time.Millisecond)
		}
	}()

	for {
		fmt.Println("runtime.NumGoroutine() :", runtime.NumGoroutine())
		time.Sleep(2 * time.Second)
	}

}
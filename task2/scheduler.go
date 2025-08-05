
package main

import (
	"fmt"
	"sync"
	"time"
)

// 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 考察点 ：协程原理、并发任务调度。

// 切片在内存中是连续的

type Task func()


type TaskResult struct{
	taskId int
	startTime time.Time
	endTime time.Time
	duration time.Duration
}

type Scheduler struct{
	tasks []Task
	results []TaskResult
	wg sync.WaitGroup
	startTime time.Time
}


func newScheduler(tasks []Task) *Scheduler{
	return &Scheduler{
		tasks: tasks,
		results: make([]TaskResult, len(tasks))
	}
}

func (s *Scheduler) executeTask(id int, task Task){
	defer s.wg.Done()
	start := time.Now()
	task()
	end := time.Now()

	s.results[id] = TaskResult{
		taskId: id,
		startTime: start,
		endTime: end,
		duration: end.Sub(start)
	}
}

func (s *Scheduler) printResult(){
	fmt.Println("------- result --------")
	for _, result := range s.results {
		fmt.Printf(
			"任务id:%d, 任务开始时间:%v, 任务结束时间:%v, 任务持续时间:%v; \n",
			result.taskId, 
			result.startTime.Format("15:04:05.000"), 
			result.endTime.Format("15:04:05.000"), 
			result.duration,
		)
	}
	fmt.Println("执行总时长：", time.Since(s.startTime))
}


func (s *Scheduler) run() {
	s.startTime = time.Now()
	for i, task := range s.tasks{
		s.wg.Add(1)
		go s.executeTask(i, task)
	}
	s.wg.Wait()
}


func main(){
	// 示例任务定义
	tasks := []Task{
		func() {
			time.Sleep(500 * time.Millisecond)
			fmt.Println("任务1完成")
		},
		func() {
			time.Sleep(300 * time.Millisecond)
			fmt.Println("任务2完成")
		},
		func() {
			time.Sleep(200 * time.Millisecond)
			fmt.Println("任务3完成")
		},
	}

	s := newScheduler(tasks)
	s.run()
	s.printResult()
}
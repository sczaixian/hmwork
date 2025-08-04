
package main

import (
	"fmt"
	"sync"
	"time"
)

// 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 考察点 ：协程原理、并发任务调度。


type Task func()


type TaskResult struct{
	taskId int
	startTime time.Time
	endTime time.time
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
}


func (s *Scheduler) run() {
	s.startTime = time.Now()
	for i, task := range s.tasks{
		s.wg.add(1)
		go s.executeTask(i, task)
	}
	s.wg.Wait()
}
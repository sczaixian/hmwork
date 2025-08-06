package main
import (
	"fmt"
	"sync"
	"sync/atomic"
)

func xxx(){
	// buf_ch := make(chan int, 20)
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go func(){
		defer (
			wg.Done()
			Close(ch)
		)
		for i := 1; i < 10; i++ {
			ch <- i
		}
	}()

	go func(){
		defer wg.Done()
		for v:= range ch {
			// print v
		}
	}()
	
	wg.Wait()
}

type SafeCounter struct {
	count int
	mtx sync.Mutex
}

func (s * SafeCounter) inc (){
	// 互斥锁
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.count++
	// 原子操作
}


func main(){

	fmt.Println("----hello world-----")
	
}
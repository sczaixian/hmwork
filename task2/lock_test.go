

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)


// 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。
// 启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ： sync.Mutex 的使用、并发数据安全。

type safeCounter struct{
	count int
	mtx sync.Mutex
}


func (sc *safeCounter) inc{
	sc.mtx.Lock()
	defer sc.mtx.Unlock()
	sc.count++
}

func demo1(){
	counter := safeCounter{count: 0}

	const c_num = 10
	const inc_num = 1000    //  不是 :=  而是 = 

	var wg sync.WaitGroup
	wg.Add(c_num)

	for n:= 0; n < c_num; n++{
		go func(){
			defer wg.Done()
			for i:=0; i < inc_num; i++{
				counter.inc()
			}
		}()
	}
	

	wg.Wait()
	fmt.Println("result: ", counter.count)
}


// 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，
// 每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ：原子操作、并发数据安全。
func demo2(){
	var counter = uint64(0)
	var wg  sync.WaitGroup
	wg.Add(10)

	for i:= 0; i < 10; i++ {
		go func(id int){
			defer wg.Done()

			for j := 0; j < 1000; j++ {
				atomic.AddUint64(&counter, 1)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println(counter)
	// finalValue := atomic.LoadUint64(&counter)
	// fmt.Printf("原子加载验证: %d\n", finalValue)
}


func main(){
	demo1()
	demo2()
}


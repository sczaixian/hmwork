package main

import (
	"fmt"
	"sync"
)

// 题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，
// 并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
// 考察点 ：通道的基本使用、协程间通信。

func demo1() {
	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		// defer是Go语言中的一个关键字，用于延迟函数的执行。
		// 它通常用于在函数结束时执行一些清理工作，如关闭文件、释放资源等。
		// defer语句会将其后的函数调用推迟到包含它的函数执行完毕时执行。
		defer wg.Done() // 协程结束时通知WaitGroup
		defer close(ch)
		for i := 0; i < 10; i++ {
			ch <- i
			fmt.Println("send:", i)
		}
		fmt.Println("finish producer!")
	}()

	go func() {
		defer wg.Done()
		for n := range ch {
			fmt.Println("recv:", n)
		}
		fmt.Println("finish consumer!")
	}()

	wg.Wait()
	fmt.Println("finish!")
}

// 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
// 考察点 ：通道的缓冲机制。

func demo2() {
	buf_ch := make(chan int, 10)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		defer close(buf_ch)

		for i := 0; i < 10; i++ {
			buf_ch <- i
			fmt.Printf(
				"send: %d,  buffer(%d/%d) \n", i, len(buf_ch), cap(buf_ch),
			)
		}
		fmt.Println("finish producer!")
	}

	go func() {
		defer wg.Done()
		for n := range buf_ch {
			fmt.Printf(
				"recv: %d,  buffer(%d/%d) \n", n, len(buf_ch), cap(buf_ch),
			)
		}
		fmt.Println("finish consumer!")
	}

	wg.Wait()
	fmt.Println("finish!")
}

// func demo3(){
// 	go func() {
// 		defer func(){
// 			wg.Done()
// 			close(buf_ch)
// 		}

// 		for i := 0; i < 10; i++ {
// 			select {
// 			case buf_ch <- i: // 正常发送
// 				fmt.Printf("send: %d\n", i)
// 			case <-time.After(1 * time.Second): // 防止永久阻塞
// 				fmt.Println("发送超时")
// 				return
// 			}
// 		}
// 	}()
// }

func main() {
	demo1()
	demo2()
}

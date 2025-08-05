
// main 函数必须在 main 包中才能作为可执行程序

package main   // 必须改为main包才能作为可执行程序

import  (
	"fmt"
	"sync"
)

// 题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，
// 在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。

func foo(input *int) {
	*input += 10
	/*
	var p *int
	var p := new(int)

	if p == nil // 空指针
	*/
}

// 实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2
func foo2(input *[]int) {  // Go中的切片已经是引用类型
	slice := *input
	for i,_ := range slice {
		slice[i] *= 2
	}
}

// 编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数
func goRoutine(){
	var wg sync.WaitGroup
	wg.Add(2)

    go func(){
		defer wg.Done()  // 每个协程结束时调用 wg.Done()
		for i := 1; i <= 10; i++{
			if i % 2 != 0{
				fmt.Println(i)
			}
		}
	}()

	go func(){
		defer wg.Done()
		for i := 2; i <= 10; i++{
			if i % 2 == 0{
				fmt.Println(i)
			}
		}
	}()

	wg.Wait() // 阻塞主协程直到所有协程完成
}


// 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 考察点 ：协程原理、并发任务调度。





func main(){
	a := 10
	foo(&a)
	fmt.Println(a)

	arr := []int{1,2,3,4,5}
	foo2(&arr)
	fmt.Println(arr)
// 	for _, v := range arr{
// 		fmt.Println(v)
// 	}
}
package main

import (
	"fmt"
)

// 题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，
// 再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
// 为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
// 考察点 ：组合的使用、方法接收者。

/*
继承：  is-a
组合：  has-a
*/

type Person struct {
	name string
	age  int
}

type Employee struct {
	Person
	employeeId int
}

func (e Employee) printInfo() {
	fmt.Printf(
		"name:%s, age:%d, employeeid:%d \n", e.name, e.age, e.employeeId,
	)
}

func main() {
	e := Employee{
		Person{
			name: "Jack",
			age:  20, // ,
		}, // ,
		employeeId: 2001,
	}
	e.printInfo()

}

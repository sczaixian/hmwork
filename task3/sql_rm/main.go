package main

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm/clause"
)


func conn_mysql() *gorm.DB,error {
	dsn := "sc:123@tcp(182.168.3.52:3306)/task3?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db, nil
}

// 题目1：使用SQL扩展库进行查询
// 假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。

type Employee struct {
	ID uint
	Name string
	Department string
	Salary float64
}

type APIEmployee struct{
	ID uint
	Name string
	Department string
	Salary float64
}

// 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
func foo1(db *gorm.DB){
	var results []APIEmployee
	db.Model(Employee{}).Where("department = ?", "技术部").Find(&results)
	if results.Error != nil{
		fmt.Println("error", results.Error)
		return
	}
	for _, res := range results {
		fmt.Println(res.ID, res.Name, res.Department, res.Salary)
	}


}
// 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
func foo2(db *gorm.DB){
	var res Employee
	db.Table("employees").Order("salary desc").First(&res)
	if res.Error != nil {
		fmt.Println("err: ", res.Error)
		return
	}
	fmt.Println(res.ID, res.Name, res.Department, res.Salary)
}


// 假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
// 定义一个 Book 结构体，包含与 books 表对应的字段。
type Book struct{
	ID uint `gorm:"primarykey;autoIncrement"`
	Title string `gorm:"type :varchar(255);not null;index"`
	Author string 
	Price float64
}

// 编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
func test(db * grom.DB){
	var results []Book
	db.Model(Book{}).Where("price > ?", 50).Order(clause.OrderByColumn{
		Clumn: clause.Clumn{Name: "price"},
		Desc: "desc",
	}).Find(&results)
	for _, res := range results {
		fmt.Println(res.ID, res.Title, res.Author, res.Price)
	}
}


func demo2(db * gorm.DB){
	books := []Book{
		{Title: "Go语言高级编程", Author: "柴树杉", Price: 89.90},
		{Title: "深入理解计算机系统", Author: "Randal E.Bryant", Price: 139.00},
		{Title: "算法导论", Author: "Thomas H.Cormen", Price: 128.00},
		{Title: "数据库系统概念", Author: "Abraham Silberschatz", Price: 99.50},
		{Title: "畅销作家写作技巧", Author: "大泽在昌", Price: 45.00},
		{Title: "Python数据分析", Author: "Wes McKinney", Price: 75.00},
		{Title: "JavaScript权威指南", Author: "David Flanagan", Price: 118.00},
		{Title: "云原生架构", Author: "畅销技术团队", Price: 65.00},
	}
	db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&books).Error; err != nil {
			return err
		}
		return nil
	})
}


func main() {
	db := conn_mysql()
	if err := db.AutoMigrate(&Employee{}); err != nil {
		panic("数据库迁移失败: " + err.Error())
	}
	employees := []Employee{
		{ID:1, Name:"张三", Department:"市场部", Salary:2000},
		{ID:2, Name:"李四", Department:"技术部", Salary:3000},
		{ID:3, Name:"王五", Department:"财务部", Salary:4000},
		{ID:4, Name:"赵六", Department:"商务部", Salary:5000},
		{ID:5, Name:"钱七", Department:"销售部", Salary:6000},
		{ID:6, Name:"孙八", Department:"策划部", Salary:7000},
		{ID:7, Name:"周九", Department:"运营部", Salary:8000},
	}
	db.Create(&employees)
	foo1(db)
	foo2(db)
	if err := db.AutoMigrate(&Book{}); err != nil {
		panic("数据库迁移失败: " + err.Error())
	}
	demo2(db)

	fmt.Println("hello world")
}

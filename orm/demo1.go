

package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main(){
	db, err := gorm.Open(mysql.Open("sc:123@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
	if err!=nil{
		panic(err)
	}
	
}

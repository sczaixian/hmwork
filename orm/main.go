
package main



import (
	"hmwork/orm/conn"
	"fmt"
)

func main(){
	db := conn.InitDB()
	fmt.Println(db)
}


/*
import (
	"hmwork/orm/test"
)

func main(){
	test.TestHello()  // 子目录
	orm.TestImportHello() // 同级目录 不用导 编译 使用  go run .
}


导包问题：
go mod init  hmwork
必须将函数名改为大写开头，不然到不进去
不支持相对路径

*/
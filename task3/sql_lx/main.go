package main

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// const FLAG = "generics" // traditional
const FLAG = "traditional"

func conn_mysql() *gorm.DB {
	dsn := "sc:123@tcp(192.168.3.52:3306)/task3?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	return db
}

type CRUD interface {
	Create(db *gorm.DB)
	Update(db *gorm.DB, conditions ...string)
	Delete(db *gorm.DB)
	InitTable(db *gorm.DB)
}

type User struct {
	gorm.Model
	Name  string
	Age   int
	Grade string
}

func (u *User) InitTable(db *gorm.DB) {
	db.AutoMigrate(u)
}

func (u *User) Create(db *gorm.DB) {
	if FLAG == "generics" {
		ctx := context.Background()
		result := gorm.WithResult()
		err := gorm.G[User](db, result).Create(ctx, u)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result.RowsAffected, u.ID)
	} else {
		result := db.Create(u)
		if result.Error != nil {
			fmt.Println(result.Error)
		}
		fmt.Println(result.RowsAffected, u.ID)
	}
}

func UserTest(db *gorm.DB) {
	//编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"
	u := &User{
		Name:  "张三",
		Age:   30,
		Grade: "三年级",
	}
	u.ID = 5 // 增加固定id 不会新插入数据
	u.InitTable(db)
	u.Create(db)

	mult_users := []User{
		{Name: "张三", Age: 30, Grade: "一年级"},
		{Name: "李四", Age: 30, Grade: "二年级"},
		{Name: "王五", Age: 30, Grade: "三年级"},
	}
	db.Create(&mult_users)

	//编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
	users := []User{}
	db.Where("Age > ?", 18).Find(&users)

	for _, user := range users {
		fmt.Println("after insert: ", user.ID, user.Name, user.Age, user.Grade)
	}

	ctx := context.Background()
	users2, err := gorm.G[User](db).Find(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("users2 ---> ", len(users2))

	//编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	gorm.G[User](db).Where("Name = ?", "张三").Update(ctx, "Grade", "四年级")

	db.Find(&users)
	for _, user := range users {
		fmt.Println("after update: ", user.ID, user.Name, user.Age, user.Grade)
	}

	db.Model(User{}).Where("Name = ?", "张三").Update("Grade", "三年级")
	db.Find(&users)
	for _, user := range users {
		fmt.Println("after rollback: ", user.ID, user.Name, user.Age, user.Grade)
	}

	s_user := &User{}
	db.Model(User{}).Where("id = ?", 3).Take(&s_user) //first last
	fmt.Println("id=3 --> : ", s_user.ID, s_user.Name, s_user.Age, s_user.Grade)
	result_user := map[string]interface{}{}
	db.Table("users").Where("id = ?", 2).Take(&result_user)
	for k, v := range result_user {
		fmt.Printf("%s: %s \t", k, v)
	}
	fmt.Println()
	db.Model(User{}).Where("id = ?", 2).Updates(User{Name: "新张三", Grade: "一年级", Age: 10})
	db.Find(&users)
	for _, user := range users {
		fmt.Println("updates: ", user.ID, user.Name, user.Age, user.Grade)
	}

	//编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
	gorm.G[User](db).Where("age < ?", 15).Delete(ctx)
	db.Find(&users)
	for _, user := range users {
		fmt.Println("delete1: ", user.ID, user.Name, user.Age, user.Grade)
	}

	db.Where("age = ?", 30).Delete(users)
	db.Find(&users)
	for _, user := range users {
		fmt.Println("delete2: ", user.ID, user.Name, user.Age, user.Grade)
	}
}

type Account struct {
	ID      uint `gorm:"primary_key"`
	Balance uint `gorm:"not null; check:balance >= 0"`
}

type Transaction struct {
	ID            uint `gorm:"primary_key"`
	FromAccountID uint `gorm:"not null; index"`
	ToAccountID   uint `gorm:"not null; index"`
	Amount        uint `gorm:"not null; check:amount > 0"`
}

func transferMoney(db *gorm.DB, from uint, to uint, amount uint) error {
	if from == to {
		return errors.New("相同账户")
	}
	if amount < 0 {
		return errors.New("金额小于0")
	}
	db.Transaction(func(tx *gorm.DB) error {
		var accounts []Account
		err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("ID IN (?, ?)", from, to).Find(&accounts).Error
		if err != nil {
			return err
		}

		var amountA, amountB *Account
		for idx := range accounts {
			if accounts[idx].ID == from {
				amountA = &accounts[idx]
			} else if accounts[idx].ID == to {
				amountB = &accounts[idx]
			}
		}

		if amountA != nil || amountB != nil {
			return errors.New("账户不对")
		}

		if amountA.Balance < amount {
			return errors.New("余额不足")
		}

		if err := tx.Model(Account{}).Where("id = ?", from).Update("Balance", gorm.Expr("blance - ?", amount)).Error; err != nil {
			return err
		}
		if err := tx.Model(Account{}).Where("id = ?", to).Update("Balance", gorm.Expr("blance + ?", amount)).Error; err != nil {
			return err
		}

		trx := Transaction{
			FromAccountID: from,
			ToAccountID:   to,
			Amount:        amount,
		}
		if err := tx.Create(&trx).Error; err != nil {
			return err
		}
		return nil
	})

	return nil
}

func TransactionTest(db *gorm.DB) {
	//db.AutoMigrate(&Transaction{}, &Account{})
	//accounts := []Account{
	//	{Balance: 100, ID: 1},
	//	{Balance: 100, ID: 2},
	//	{Balance: 0, ID: 3},
	//	{Balance: 0, ID: 4},
	//}
	//db.Create(&accounts)
	err := transferMoney(db, 1, 2, 100)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("1111 success!!")

	//err2 := transferMoney(db, 3, 4, 100)
	//if err2 != nil {
	//	fmt.Println(err2)
	//	return
	//}
	//fmt.Println("2222 success!!")
}

func main() {
	db := conn_mysql()
	//UserTest(db)

	TransactionTest(db)

	fmt.Println("---------")
}

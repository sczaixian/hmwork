package crud

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"hmwork/orm/util"
	"time"
)

type User struct {
	gorm.Model
	Name     string
	Age      int
	Birthday time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Name = u.Name + "_123"
	return
}

func CreateDemo(db *gorm.DB) {
	db.AutoMigrate(&User{})

	user := User{
		Name:     "Jack",
		Age:      18,
		Birthday: time.Now(),
	}

	users := []*User{
		{Name: "Jinzhu", Age: 18, Birthday: time.Now()},
		{Name: "Jackson", Age: 19, Birthday: time.Now()},
	}

	// 创建单条记录
	//result := db.Create(&user)
	//fmt.Println(result.RowsAffected)

	ctx := context.Background()
	err := gorm.G[User](db).Create(ctx, &user)
	if err != nil {
		fmt.Println(err)
	}

	result := gorm.WithResult()
	err1 := gorm.G[User](db, result).Create(ctx, &user)
	if err != nil {
		fmt.Println(err1)
	}

	new_result := db.Create(&users)
	fmt.Println(new_result.RowsAffected, new_result.Error)
}

func SelectDemo(db *gorm.DB) {
	ctx := context.Background()
	user, err := gorm.G[User](db).Find(ctx) // select * from users order by id limit 1;
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
	/*
		.Take(ctx)  SELECT * FROM users LIMIT 1;
		.Last(ctx)  SELECT * FROM users ORDER BY id DESC LIMIT 1;
	*/

	/*
		db.First(&user)
		db.Take(&user)
		db.Last(&user)
	*/

	// 根据主键检索
	//SELECT * FROM users WHERE id = 10;
	user2, err := gorm.G[User](db).Where("id=?", 10).First(ctx)
	util.Print(user2, err)

	//SELECT * FROM users WHERE id = 10;
	user3, err := gorm.G[User](db).Where("id=?", 20).First(ctx)
	util.Print(user3, err)

	user4, err := gorm.G[User](db).Where("id IN ?", []int{1, 2, 3}).Find(ctx)
	util.Print(user4, err)

	/*
		db.First(&user, 10)
		db.First(&user, "10")
		db.Find(&users, []int{1,2,4})
		db.Find(&users)  // select * from users;

		name <> ?    name = ?   name in ?   name like ? "%ch%"
		name = ? and age >= ?


		var user = User{ID: 10}
		db.Where("id = ?", 20).First(&user)
		select * from users where id = 10 and id = 20 order by id asc limit 1;


		db.Where(&User{Name:"jack", Age: 20,}).First(&user)
		select * from users where name="jack" and age = 20 order by id limit 1

		db.Where(map[string] interface{}{"name":"jack", "age":20}).Find(&users)
		select * from users where name="jack" and age = 20;

		db.Where([]int{1,2,3}).Find(&users)
		select * from users where id in (1,2,3)

		db.Where(&User{name:"jack"}, "name", "age").Find(&users)
		select * from users where name="jack" and age = 0  //  age 没有明确指定，默认0
		db.where(&User{name:"jack"}, "age").Find(&users)
		select * from users where age = 0;

		db.Not("name = ?", "jack).Find(&users)
		select * from users where name != 'jack'

		db.Where("role = ?", "admin").Or("role = ?", "aa").Find(&users)
		select * from users where role = "admin" or role = "aa";

		db.Select("name", "age").Find(&users)
		select name, age form users;

	
	*/

}

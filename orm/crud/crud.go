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

	/*  db.table, db.model
		model 基于模型自动推断 表名，字段映射和关系，链式查询条件可以用结构体字段名，支持 hook
		
		table 接收字符串表名， 需要使用数据库原生列名，不触发hook 没有字段映射， 结果类型：map、非模型结构体、动态类型

		db.Model的自动映射略增开销，但简化代码；适合模型驱动、结构化场景，强调约定优于配置
		db.Table更轻量，适合高频动态操作；提供最大灵活性，适合动态表、原生SQL等复杂需求
	*/

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

		Scan 直接映射原生 SQL 结果到结构体，不依赖模型定义；Find 则需完整模型结构
		type result struct{
			Name string
			Email string
		}
		db.Model(&User{}).Select("users.Name, emails.Email").Joins("left join emails on users.user_id = emails.user_id").Scan(&result{})
		select users.Name,emails.Email from users left join emails on user.user_id=emails.user_id;

		rows, err:= db.Table("users").select("users.Name, emails.Email").Joins("left join emails on users.user_id=emails.user_id").Rows()
		for rows.Next(){
			...
		}


		type User struct {
		    Id int
			Age int
		}
		
		type Order struct {
		    UserId int
			FinishedAt *time.Time
		}
		query := db.Talbe("order").Select("max(orders.finished_at) as latest")
								   .Joins("left join user user on orders.user_id=users.user_id")
								   .Where("users.age > ?", 18).Group("orders.user_id")

								   
		db.Model(&Order{}).Select("orders.user_id, orders.finished_at")
						  .Joins("join (?) q on orders.finished_at=q.latest")
						  .Scan(&rusult{})

		SELECT `order`.`user_id`,`order`.`finished_at` FROM `order` 
		join (
		SELECT MAX(order.finished_at) as latest FROM `order` 
		left join user user 
		on order.user_id = user.id 
		WHERE user.age > 18 GROUP BY `order`.`user_id`
		) q on order.finished_at = q.latest

	*/

	/*
	type User struct {
		// 。。。很多字段
	}
	
	type APIUser struct{
		Name string
		Age int
	}
	db.Model(&User{}).Limit(10).Find(&APIUser{})
	select Name, Age from users limit 10;
	
	*/

	// ​​QueryFields 模式​​ 是一种优化查询性能的重要机制，
	// 它通过显式指定查询字段而非使用 SELECT *，减少数据传输量、避免敏感字段泄露，并提升数据库执行效率
	/**
	db, err := grom.Open(sqlite.Open("gorm.db"), &gorm.Config{
		QueryFields: true, // 全局生效
	})
	
	db.Session(&gorm.Session(QueryFields: true)).Find(&users)
	*/



	// 锁
	/*
	select for update 
	db.Clauses(clause.Locking{Strength: "update"}).Find(&users)
	select * from users for update


	select for share of
	db.Clauses(clause.Locking{
		Strength: "share",
		Table: clause.Table{Name: clause.CurrentTable}  // 指定锁定的表，在join 多个表时
		Options: "nowait/skip locked" // nowait: 尝试获取锁，没成功报错；  skip locked 跳过所有加锁的行
	}).Find(&users)
	select * from users for share users
	*/

	// 使用 (?) 嵌入子查询

	// 多列 in  
	/*
	db.Where(
		"(name, age, email) in ?", [][]interface{
			{"name1", 10, "email1"},
			{"name2", 11, "email2"},
			{"name3", 12, "email3"},
		}
	)
	WHERE (name, age, role) IN (("jinzhu", 18, "admin"), ("jinzhu 2", 19, "user"))
	*/

	/*
	命名参数： sql.Named("arg1","arg2")  map[string]interface{}{"a","b"}
	db.Where("name1=@name and name2=@name", sql.Named("a","b"))

	db.Where("name1=@name and name2=@name", map[string]interface{}{"a","b"})
	*/

	/*
	result := map[string]interface{}{}
	db.Model(&User{}).First(&result, "id = ?", 20)
	select * from users where id = 20 limit 1

	var results []map[string]interface{}{}
	db.Table("users").Find(&result, "id > ?", 1)
	select * from users where id > 1
	
	*/

	/*
		分批查询 FindInBatches
		result := db.Where("id > ?", 20).FindInBatches(&results, 100, func (ctx * gorm.DB, batch int) error {
			for _ result := range results {
				// 对批中的每条记录进行操作
			}
			tx.Save(&results)
			return nil  // error
		})
	
	*/


	/*	hook 在查询声明周期中触发
		BeforeSave, BeforeUpdate, AfterSave, AfterUpdate

		func (u * User)AfterFind(tx * grom.DB) (err error) {
			//  ....
			return nil
		}
		// 当user 查询时会自动触发
	*/

	// Pluck 单列切片, 多列用 Scan，Find
	/*
		var names []string
		db.Model(&User{}).Where("id > ?", 20).Distinct().Pluck("name", &res)
		db.Table("tablename").Pluck("name", &names)
	*/


	/*  Scope 将常用的查询条件定义为可重用的方法

		func AmountGreaterThan1000 (db * gorm.DB) (*gorm.DB) {
			return db.Where("amount > ?", 1000)
		}

		func OrderStatus(status []string) func(db * gorm.DB) *gorm.DB {
			return func (db * gorm.DB) *gorm.DB{
				return db.Where("status in (?)", status)
			}
		}

		db.Scopes(AmountGreaterThan1000, ...其他scope)
		db.Scopes(AmountGreaterThan1000, OrderStatus([]string{"xx","xx"})).Find(&orders)
	*/



	/*  更新
		.Update("name", "hello")  // 单列
		ctx := context.BackGround()  //  更新多列
		err := gorm.G[&User](db).Where("id = ?", 20).Updates(ctx, User{Name:"xx",Age:20,Email:"xx"})
		update users set name="", age=20, email="xx",updated_at="xx:xx:xx" where id=20
		err := db.Model(&users).Updates(map[string]interface{}{"name":"","age":"","email":""})


		UpdateColumn, UpdateColumns 不使用 hook 和 更新时间戳
	*/


	/*	Statement 测试
		// 新建会话模式
		stmt := db.Session(&Session{DryRun: true}).First(&user, 1).Statement
		stmt.SQL.String() //=> SELECT * FROM `users` WHERE `id` = $1 ORDER BY `id`
		stmt.Vars         //=> []interface{}{1}
	
	*/
}

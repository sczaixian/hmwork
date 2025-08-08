package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	//"gorm.io/gorm/clause"
)

func conn_mysql() *gorm.DB {
	dsn := "sc:123@tcp(192.168.3.52:3306)/task3?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return db
}

type PostStatus struct {
}

// 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
// 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章），
//
//	Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
type User struct {
	gorm.Model
	Name      string
	Password  string
	PostCount uint
	Post      []Post
}

type Post struct {
	gorm.Model
	Title        string
	Context      string
	UserID       uint
	CommentCount uint
	// Status PostStatus
	Status  string
	Comment []Comment
}

type Comment struct {
	gorm.Model
	Context string
	UserID  uint
	PostID  uint
}

// 基于上述博客系统的模型定义。
// 要求 ：
// 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
func demo1(db *gorm.DB, userId uint) ([]Post, error) {
	posts := []Post{}
	err := db.Model(&Post{}).Where("user_id = ?", userId).Preload("Comments", func(db *gorm.DB) *gorm.DB {
		return db.Order("comments.create_at desc")
	}).Find(&posts)
	// Association("Comments").
	if err.Error != nil {
		return nil, err.Error
	}
	return posts, nil
}

// 编写Go代码，使用Gorm查询评论数量最多的文章信息。
func demo2(db *gorm.DB) (Post, error) {
	post := Post{}
	subQuery := db.Model(&Comment{}).Select("PostId, count(*) as num").Group("PostID")
	err := db.Model(&Post{}).Select("posts.*, a.num").Joins("(?) as a left join posts.id=a.post_id", subQuery).Order("a.num desc").First(&post)
	if err.Error != nil {
		if err.Error == gorm.ErrRecordNotFound {
			return Post{}, fmt.Errorf("没找到任何文章")
		}
		return Post{}, err.Error
	}
	return post, nil
}

// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
func (p *Post) AfterCreate(tx *gorm.DB) error {
	/*
		UpdateColumn
			跳过钩子(Hooks)：UpdateColumn会跳过所有的回调方法（BeforeSave, BeforeUpdate, AfterSave, AfterUpdate 等）
			不更新更新时间：不会自动更新 updated_at时间戳
			直接SQL更新：直接生成 SQL 语句执行更新，不进行字段值的验证
			不加载关联：不会处理关联关系

				需要高性能更新，跳过所有回调
				不需要更新 updated_at时间戳
				确定数据有效性，不需要验证

		Update
			执行钩子：会触发 BeforeSave, BeforeUpdate, AfterSave, AfterUpdate 等回调方法
			更新更新时间：会自动更新 updated_at时间戳
			字段验证：会验证字段值
			处理关联：可以处理关联关系

				需要完整的模型生命周期
				需要自动更新时间戳
				需要字段验证
	*/
	if err := tx.Model(&User{}).Where("id = ?", p.UserID).Update("PostCount", gorm.Expr("post_count + ", 1)).Error; err != nil {
		return err
	}
	return nil
}

// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
func (c *Comment) AfterDelete(tx *gorm.DB) error {
	var comment_count int64
	if err := tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&comment_count).Error; err != nil {
		//
	}
	if comment_count == 0 {
		if err := tx.Model(&Post{}).Where("id = ?", c.PostID).Update("status", "无评论").Error; err != nil {
			//
		}
	}

	return nil
}

func main() {
	db = conn_mysql()
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
	//users := []User{
	//	{ID: 1, Name: "", Password: "", }
	//}
}

package create

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint           // Standard field for the primary key
	Name         string         // A regular string field
	Email        *string        // A pointer to a string, allowing for null values
	Age          uint8          // An unsigned 8-bit integer
	Birthday     *time.Time     // A pointer to time.Time, can be null
	MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
	ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
	CreatedAt    time.Time      // Automatically managed by GORM for creation time
	UpdatedAt    time.Time      // Automatically managed by GORM for update time
	ignored      string         // fields that aren't exported are ignored
}

/*
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` longtext COLLATE utf8mb4_unicode_ci,
  `email` longtext COLLATE utf8mb4_unicode_ci,
  `age` tinyint unsigned DEFAULT NULL,
  `birthday` datetime(3) DEFAULT NULL,
  `member_number` longtext COLLATE utf8mb4_unicode_ci,
  `activated_at` datetime(3) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci


 CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` longtext COLLATE utf8mb4_unicode_ci,
  `email` longtext COLLATE utf8mb4_unicode_ci,
  `age` tinyint unsigned DEFAULT NULL,
  `birthday` datetime(3) DEFAULT NULL,
  `member_number` longtext COLLATE utf8mb4_unicode_ci,
  `activated_at` datetime(3) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
*/

type Member struct {
	gorm.Model
	Name string
	Age  uint8
}

/*
CREATE TABLE `members` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext COLLATE utf8mb4_unicode_ci,
  `age` tinyint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_members_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
*/

type Author struct {
	Name  string
	Email string
}

type Blog struct {
	Author
	ID      uint
	Upvotes int64
}

/*
CREATE TABLE `blogs` (
  `name` longtext COLLATE utf8mb4_unicode_ci,
  `email` longtext COLLATE utf8mb4_unicode_ci,
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `upvotes` bigint DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
*/

type Blog2 struct {
	ID      uint
	Author  Author `gorm:"embedded;embedded_prefix:author_"`   // 给 这张表中 Author中的属性字段前 加上前缀
	Upvotes int32
}

func Run(db *gorm.DB) {
	db.AutoMigrate(&User{}) // 建 user 表
	db.AutoMigrate(&Member{})
	db.AutoMigrate(&Blog{})
	db.AutoMigrate(&Blog2{})

	user := &User{}
	user.MemberNumber.Valid = true
	db.Create(user)

	// create
	mem := &Member{}
	db.Create(&mem)
	fmt.Println(mem.ID)
	db.Delete(&Member{}, 1)
	fmt.Println("--------------")
}

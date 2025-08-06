

package orm

import (
	"fmt"
  "context"
  "gorm.io/driver/sqlite"
  "gorm.io/gorm"
)

type Product struct {
  gorm.Model
  Code  string
  Price uint
}

func test() {
  db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  ctx := context.Background()

//   // Migrate the schema
//   db.AutoMigrate(&Product{})

//   // Create
//   err = gorm.G[Product](db).Create(ctx, &Product{Code: "D42", Price: 100})
//   fmt.Println(err)


  // Read
  product, err := gorm.G[Product](db).Where("id = ?", 1).First(ctx) // find product with integer primary key
  fmt.Println(product, err)
  products, err := gorm.G[Product](db).Where("code = ?", "D42").Find(ctx) // find product with code D42
	fmt.Println(products, err)
//   // Update - update product's price to 200
//   err = gorm.G[Product](db).Where("id = ?", product.ID).Update(ctx, "Price", 200)
//   // Update - update multiple fields
//   err = gorm.G[Product](db).Where("id = ?", product.ID).Updates(ctx, map[string]interface{}{"Price": 200, "Code": "F42"})

//   // Delete - delete product
//   err = gorm.G[Product](db).Where("id = ?", product.ID).Delete(ctx)
}
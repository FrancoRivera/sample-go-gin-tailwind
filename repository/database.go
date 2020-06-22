// previous to run
// go get -u github.com/lib/pq

package repository

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	dbname = "houses"
)

var Database *gorm.DB

func SetupDatabase() {
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASSWORD")

	var err error

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	Database, err = gorm.Open("postgres", psqlInfo)

	if err != nil {
		panic("failed to connect database")
	}
	//defer Database.Close()

	// Migrate the schema
	// db.AutoMigrate(&models.House{})

	// // Create
	// db.Create(&Product{Code: "L1212", Price: 1000})
	//
	// // Read
	// var product Product
	// db.First(&product, 1) // find product with id 1
	// db.First(&product, "code = ?", "L1212") // find product with code l1212
	//
	// // Update - update product's price to 2000
	// db.Model(&product).Update("Price", 2000)
	//
	// // Delete - delete product
	// db.Delete(&product)
}

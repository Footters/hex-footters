package mysqldb

import (
	"os"

	"github.com/Footters/hex-footters/pkg/media"
	"github.com/jinzhu/gorm"
)

// NewMySQLConnection func
func NewMySQLConnection() *gorm.DB {

	var db *gorm.DB
	var err error

	db, err = gorm.Open("mysql", os.Getenv("MYSQL_CONNECTION"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&media.Content{})

	return db
}

package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //yb
	"qip.io/q/pkg/qipp/model"
)

var (
	DbConn *gorm.DB
)

func InitDb() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	db := os.Getenv("DB_DATABASE")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")

	prosgret_conname := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v", host, port, user, db, password)

	var err error
	DbConn, err = gorm.Open("postgres", prosgret_conname)
	if err != nil {
		panic("Failed to connect to database!")
	}
	DbConn.AutoMigrate(&model.Qipp{})
	fmt.Println("Database initialized")
}

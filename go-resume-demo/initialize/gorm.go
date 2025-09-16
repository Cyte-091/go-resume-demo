package initialize

import (
	"github.com/Cyte-091/go-resume-demo/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitGorm() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/resume_demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("mysql error:", err)
	}
	db.AutoMigrate(&model.Article{})
	DB = db
	log.Println("AutoMigrate success")
}

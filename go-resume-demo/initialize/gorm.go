package initialize

import (
	"github.com/Cyte-091/go-resume-demo/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func InitGorm() {
	// 从环境变量获取DSN，如果不存在则使用默认值
	dsn := os.Getenv("DB_SOURCE")
	if dsn == "" {
		// 默认使用MySQL容器的服务名和端口
		dsn = "root:123456@tcp(127.0.0.1:3306)/resume_demo?charset=utf8mb4&parseTime=True&loc=Local"
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("mysql errno:", err)
	}
	db.AutoMigrate(&model.Article{})
	DB = db
	log.Println("AutoMigrate success")
}

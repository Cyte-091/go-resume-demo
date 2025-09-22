package initialize

import (
	"github.com/Cyte-091/go-resume-demo/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func InitGorm() {
	dsn := os.Getenv("DB_SOURCE")
	if dsn == "" {
		dsn = "root:123456@tcp(mysql:3306)/resume_demo?charset=utf8mb4&parseTime=True&loc=Local"
	}

	var db *gorm.DB
	var err error

	// 重试逻辑（最多尝试30次，每次间隔2秒）
	maxRetries := 30
	for i := 1; i <= maxRetries; i++ {
		log.Printf("尝试连接数据库 (第 %d/%d 次)...", i, maxRetries)

		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Printf("数据库连接失败: %v", err)
			time.Sleep(2 * time.Second)
			continue
		}

		// 测试连接
		sqlDB, err := db.DB()
		if err != nil {
			log.Printf("获取数据库实例失败: %v", err)
			time.Sleep(2 * time.Second)
			continue
		}

		err = sqlDB.Ping()
		if err != nil {
			log.Printf("数据库Ping失败: %v", err)
			sqlDB.Close()
			time.Sleep(2 * time.Second)
			continue
		}

		log.Println("✅ 数据库连接成功!")
		break
	}

	if err != nil {
		log.Fatalf("❌ 无法连接数据库: %v", err)
	}

	db.AutoMigrate(&model.Article{})
	DB = db
	log.Println("AutoMigrate success")
}

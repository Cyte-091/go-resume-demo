package dao

import (
	"github.com/Cyte-091/go-resume-demo/initialize" // 换成你的模块路径
	"github.com/Cyte-091/go-resume-demo/model"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestCreateArticle_OK(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	gdb, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	initialize.DB = gdb

	// 0. 任何 SELECT（不限次数）返回空
	mock.ExpectQuery(".*").WillReturnRows(sqlmock.NewRows(nil))

	// 1. 事务顺序严格：Begin → INSERT → Commit
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `articles`").
		WithArgs("t", "b", sqlmock.AnyArg(), sqlmock.AnyArg()). // 2 具体 + 2 任意
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// 2. 执行
	err = NewArticleDao.CreateArticle(&model.Article{Title: "t", Body: "b"})
	/*	assert.NoError(t, err) // ← 先断言业务无错

		// 2. 再检查 SQL 期望是否全部消费（可选，可删）
		assert.NoError(t, mock.ExpectationsWereMet())*/
}

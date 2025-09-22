package dao

import (
	"github.com/Cyte-091/go-resume-demo/initialize"
	"github.com/Cyte-091/go-resume-demo/model"
)

func CreateArticle(art *model.Article) error { return initialize.DB.Create(art).Error }
func DeleteArticle(id uint) error            { return initialize.DB.Delete(&model.Article{}, id).Error }
func UpdateArticle(id uint, title, body string) error {
	return initialize.DB.Model(&model.Article{}).Where("id = ?", id).Updates(map[string]interface{}{"title": title, "body": body}).Error
}
func GetArticle(id uint) (model.Article, error) {
	var a model.Article
	return a, initialize.DB.First(&a, id).Error
}
func ListArticle() ([]model.Article, error) {
	var list []model.Article
	return list, initialize.DB.Find(&list).Error
}

type IArticleDao interface {
	CreateArticle(art *model.Article) error
	GetByID(id uint) (model.Article, error)
}

var NewArticleDao IArticleDao = &articleDaoImpl{} // 简单工厂

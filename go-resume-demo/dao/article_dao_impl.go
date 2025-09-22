package dao

import (
	"github.com/Cyte-091/go-resume-demo/initialize"
	"github.com/Cyte-091/go-resume-demo/model"
)

type articleDaoImpl struct{}

func (d *articleDaoImpl) CreateArticle(art *model.Article) error {
	return initialize.DB.Create(art).Error
}
func (d *articleDaoImpl) GetByID(id uint) (model.Article, error) {
	var art model.Article
	return art, initialize.DB.First(&art, id).Error
}

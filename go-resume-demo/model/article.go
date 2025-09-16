package model

type Article struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Title string `json:"title" gorm:"type:varchar(100)"`
	Body  string `json:"body" gorm:"type:text"`
}

func (Article) TableName() string { return "articles" }

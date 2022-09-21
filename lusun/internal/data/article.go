package data

import (
	"gorm.io/gorm"
	"lusun/internal/biz"
	"lusun/internal/data/po"
)

type Article struct {
	Db *gorm.DB
}

func NewArticle(repo *MysqlRepo) biz.ArticleRepo {
	return &Article{Db: repo.Db}
}

func (a *Article) Detail(id int64, article *po.ZbpPost) error {
	return a.Db.Table(article.TableName()).Where("log_ID = ?", id).Limit(1).Find(article).Error
}

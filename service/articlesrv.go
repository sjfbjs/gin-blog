package service

import (
	"database/sql"
	"gin-blog/model"
	"gin-blog/model/param"
	"time"
)

type articleService struct {
}

var Article = &articleService{}
var articleDAO = &model.Article{}

func (_ *articleService) GetAll() *[]model.Article {
	return articleDAO.SelectAll()
}

func (_ *articleService) GetBySlugOrId(slugOrId string) *model.Article {
	return articleDAO.SelectBySlugOrId(slugOrId)
}

func (_ *articleService) SaveArticle(param *param.ArticleParam) bool {
	alias := sql.NullString{}
	if "" == param.Alias {
		alias.Valid = false
		alias.String = ""
	} else {
		alias.Valid = true
		alias.String = param.Alias
	}
	article := model.Article{
		Id:      param.Id,
		Title:   param.Title,
		Content: param.Content,
		Alias:   alias,
		Created: time.Now().Unix(),
	}
	var err error
	if 0 != param.Id {
		_, err = articleDAO.UpdateById(&article)
	} else {
		_, err = articleDAO.InsertOne(&article)
	}
	if nil == err {
		return true
	}
	return false
}

func (_ *articleService) DeleteArticle(id int) bool {
	return articleDAO.DeleteById(id)
}

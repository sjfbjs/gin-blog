package service

import "gin-blog/model"

type articleService struct {
}

var Article = &articleService{}
var articleDAO model.Article

func (_ *articleService) GetAll() []model.Article {
	return articleDAO.SelectAll()
}

func (_ *articleService) GetBySlugOrId(slugOrId string) model.Article {
	return articleDAO.SelectBySlugOrId(slugOrId)
}
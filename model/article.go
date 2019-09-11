package model

import (
	"log"
	"strconv"
)

type Article struct {
	Id      uint   `db:"id"`
	Title   string `db:"title"`
	Content string `db:"content"`
	Alias   string `db:"alias"`
	Created uint   `db:"created"`
}

func (_ *Article) SelectAll() *[]Article {
	articles := []Article{}
	err := DB.Select(&articles, "SELECT * FROM `article`")
	if nil != err {
		log.Println("get archives failed: ", err.Error())
		return nil
	}
	return &articles
}

func (_ *Article) SelectBySlugOrId(slugOrId string) *Article {
	article := Article{}
	id, _ := strconv.Atoi(slugOrId)
	err := DB.Get(&article, "SELECT * FROM `article` WHERE alias = ? OR id = ?", slugOrId, id)
	if nil != err {
		log.Println("get archive failed: ", err.Error())
		return nil
	}
	return &article
}

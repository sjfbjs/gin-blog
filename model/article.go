package model

import (
	"database/sql"
	"log"
	"strconv"
)

type Article struct {
	Id      uint           `db:"id"`
	Title   string         `db:"title"`
	Content string         `db:"content"`
	Alias   sql.NullString `db:"alias"`
	Created int64          `db:"created"`
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

func (_ *Article) InsertOne(article *Article) (int64, error) {
	result := DB.MustExec("INSERT INTO `article` (title, content, alias, created) VALUES (?, ?, ?, ?)",
		article.Title, article.Content, article.Alias, article.Created)
	return result.LastInsertId()
}

func (_ *Article) UpdateById(article *Article) (int64, error) {
	result := DB.MustExec("UPDATE `article` SET title = ?, content = ?, alias = ? WHERE id = ?",
		article.Title, article.Content, article.Alias, article.Id)
	return result.RowsAffected()
}

func (_ *Article) DeleteById(id int) bool {
	result := DB.MustExec("DELETE FROM `article` WHERE id = ?", id)
	_, err := result.RowsAffected()
	if nil == err {
		return true
	}
	return false
}

package model

import "log"

type Config struct {
	Id    uint   `db:"id"`
	Name  string `db:"name"`
	Value string `db:"value"`
}

func (_ *Config) SelectTitle() string {
	title := Config{}
	err := DB.Get(&title, "SELECT * FROM `config` WHERE name = ?", "title")
	if nil != err {
		log.Println("get title failed: ", err.Error())
		return ""
	}
	return title.Value
}

func (_ *Config) SelectKeyword() string {
	keyword := Config{}
	err := DB.Get(&keyword, "SELECT * FROM `config` WHERE name = ?", "keywords")
	if nil != err {
		log.Println("get keyword failed: ", err.Error())
		return ""
	}
	return keyword.Value
}
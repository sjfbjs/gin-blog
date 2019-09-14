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

func (_ *Config) UpdateTitle(title string) (int64, error) {
	result := DB.MustExec("UPDATE `config` SET value = ? WHERE name = \"title\"", title)
	return result.RowsAffected()
}

func (_ *Config) UpdateKeywords(keywords string) (int64, error) {
	result := DB.MustExec("UPDATE `config` SET value = ? WHERE name = \"keywords\"", keywords)
	return result.RowsAffected()
}

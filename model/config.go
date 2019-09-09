package model

type Config struct {
	Id    uint   `db:"id"`
	Name  string `db:"name"`
	Value string `db:"value"`
}

func (_ *Config) SelectTitle() string {
	title := Config{}
	DB.Get(&title, "SELECT * FROM `config` WHERE name = ?", "title")
	return title.Value
}

func (_ *Config) SelectKeyword() string {
	keyword := Config{}
	DB.Get(&keyword, "SELECT * FROM `config` WHERE name = ?", "keywords")
	return keyword.Value
}
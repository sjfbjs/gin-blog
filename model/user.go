package model

import "log"

type User struct {
	Id       uint   `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

func (_ *User) Auth(username, password string) *User {
	user := User{}
	err := DB.Get(&user, "SELECT * FROM `user` WHERE username = ? and password = ?", username, password)
	if nil != err {
		log.Println("auth user failed: " + err.Error())
		return nil
	}
	return &user
}

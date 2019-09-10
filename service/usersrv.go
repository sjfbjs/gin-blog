package service

import (
	"crypto/md5"
	"fmt"
	"gin-blog/model"
)

type userService struct {
}

var User = &userService{}
var userDAO = &model.User{}

func (_ *userService) Login(username, password string) *model.User {
	return userDAO.Auth(username, fmt.Sprintf("%x", md5.Sum([]byte(password))))
}
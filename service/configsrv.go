package service

import "gin-blog/model"

type configService struct {
}

var Config = &configService{}
var configDAO = &model.Config{}

func (_ *configService) GetTitle() string {
	return configDAO.SelectTitle()
}

func (_ *configService) GetKeyword() string {
	return configDAO.SelectKeyword()
}
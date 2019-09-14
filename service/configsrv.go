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

func (_ *configService) UpdateTitle(title string) bool {
	_, err := configDAO.UpdateTitle(title)
	if nil == err {
		return true
	}
	return false
}

func (_ *configService) UpdateKeyword(keywords string) bool {
	_, err := configDAO.UpdateKeywords(keywords)
	if nil == err {
		return true
	}
	return false
}

package service

import (
	"gin-blog/model"
	"gin-blog/model/param"
)

type keyService struct {
}

var Key = &keyService{}
var keyDAO = &model.Key{}

func (_ *keyService) GetAll() *[]model.Key {
	return keyDAO.SelectAll()
}

func (_ *keyService) GetByIp(ip string) *model.Key {
	return keyDAO.SelectByIp(ip)
}

func (_ *keyService) SaveKey(param *param.KeyParam) bool {
	//alias := sql.NullString{}
	//if "" == param.Alias {
	//	alias.Valid = false
	//	alias.String = ""
	//} else {
	//	alias.Valid = true
	//	alias.String = param.Alias
	//}
	key := model.Key{
		Ip:  param.Ip,
		Key: param.Key,
	}
	var err error
	if "" != param.Ip {
		_, err = keyDAO.UpdateKey(&key)
	} else {
		_, err = keyDAO.InsertOne(&key)
	}
	if nil == err {
		return true
	}
	return false
}

func (_ *keyService) DeleteKey(ip string) bool {
	return keyDAO.DeleteByIp(ip)
}

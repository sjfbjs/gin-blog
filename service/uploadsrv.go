package service

import (
	"gin-blog/model"
)

type uploadservice struct {
}

var Upload = &uploadservice{}
var uploadDAO = &model.Upload{}

func (_ *uploadservice) GetAll() *[]model.Upload {
	return uploadDAO.SelectAll()
}

func (_ *uploadservice) GetByFilename(fname string) *model.Upload {
	return uploadDAO.SelectByFilename(fname)
}

func (_ *uploadservice) SaveFile(upload *model.Upload) bool {
	//upload := model.Upload{
	//	Filename:  param.Filename,
	//	Created: param.Created,
	//}
	var err error
	hisup := uploadDAO.SelectByFilename(upload.Filename)
	if hisup != nil {
		return false
	} else {
		_, err = uploadDAO.InsertOne(upload)
	}
	if nil == err {
		return true
	}
	return false
}

func (_ *uploadservice) DeleteByFilename(fname string) bool {
	return uploadDAO.DeleteByFilename(fname)
}

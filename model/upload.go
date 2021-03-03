package model

import (
	//"database/sql"
	"fmt"
	"log"
)

type Upload struct {
	Filename string `db:"filename"`
	Created  string `db:"created"`
}

func (_ *Upload) SelectAll() *[]Upload {
	ups := []Upload{}
	err := DB.Select(&ups, "SELECT * FROM `Uploadmanage`")
	if nil != err {
		log.Println("get Upload failed: ", err.Error())
		return nil
	}
	return &ups
}

func (_ *Upload) SelectByFilename(fname string) *Upload {
	up := Upload{}
	err := DB.Get(&up, "SELECT * FROM `upload` WHERE filename = ?", fname)
	if nil != err {
		log.Println("get Upload failed: ", err.Error())
		return nil
	}
	return &up
}

func (_ *Upload) InsertOne(Upload *Upload) (int64, error) {
	fmt.Println("Upload:", Upload)
	result := DB.MustExec("INSERT INTO `Uploadmanage` (filename) VALUES (?)",
		Upload.Filename)
	return result.LastInsertId()
}

//func (_ *Upload) UpdateUpload(Upload *Upload) (int64, error) {
//	result := DB.MustExec("UPDATE `Uploadmanage` SET filename = ? WHERE ip = ?",
//		Upload.Filename)
//	return result.RowsAffected()
//}

func (_ *Upload) DeleteByFilename(fname string) bool {
	result := DB.MustExec("DELETE FROM `upload` WHERE filename = ?", fname)
	_, err := result.RowsAffected()
	if nil == err {
		return true
	}
	return false
}

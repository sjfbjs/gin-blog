package model

import (
	//"database/sql"
	"fmt"
	"log"
)

type Key struct {
	Ip  string `db:"ip"`
	Key string `db:"keycontent"`
}

func (_ *Key) SelectAll() *[]Key {
	keys := []Key{}
	err := DB.Select(&keys, "SELECT * FROM `keymanage`")
	if nil != err {
		log.Println("get key failed: ", err.Error())
		return nil
	}
	return &keys
}

func (_ *Key) SelectByIp(ip string) *Key {
	key := Key{}
	err := DB.Get(&key, "SELECT * FROM `keymanage` WHERE ip = ?", ip)
	if nil != err {
		log.Println("get key failed: ", err.Error())
		return nil
	}
	return &key
}

func (_ *Key) InsertOne(key *Key) (int64, error) {
	fmt.Println("key:", key)
	result := DB.MustExec("INSERT INTO `keymanage` (ip,keycontent) VALUES (?, ?)",
		key.Ip, key.Key)
	return result.LastInsertId()
}

func (_ *Key) UpdateKey(key *Key) (int64, error) {
	result := DB.MustExec("UPDATE `keymanage` SET keycontent = ? WHERE ip = ?",
		key.Key, key.Ip)
	return result.RowsAffected()
}

func (_ *Key) DeleteByIp(ip string) bool {
	result := DB.MustExec("DELETE FROM `keymanage` WHERE ip = ?", ip)
	_, err := result.RowsAffected()
	if nil == err {
		return true
	}
	return false
}

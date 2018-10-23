package models

import (
	"time"
)

type EquiRoom struct {
	Model 
	RoomId int64 `json:"roomid"`
	EquipName string `json:"equipname"`
	EquiNum int64 `json:"EquiNum"`
}

func CreateEquiRoom(rid int64,equipname string,num int64) error {
	db := GetSelfDB()
	var eq EquiRoom
	eq.RoomId = rid
	eq.EquipName = equipname
	eq.EquiNum = num
	eq.CreateAt = time.Now()
	return db.Create(&eq).Error
}

func DeleteEquiRoom(rid int64)error{
	db := GetSelfDB()
	er := EquiRoom{}
	er.RoomId = rid
	return db.Where("room_id = ?",er.RoomId).Delete(&er).Error
}
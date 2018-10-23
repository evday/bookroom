package models

import (
	"time"
)

type Room struct {
	Model
	Name string `json:"roomname" gorm:"column:roomname;not null"`  //会议室名称
	Location string `json:"location" gorm:"column:location;not null"`//位置
	Capacity string `json:"capacity" gorm:"column:capacity"` //可容纳人数
	State bool `json:"state" gorm:"column:state"`			//开放状态
	MeetRoom []Record `gorm:"ForeignKey:RoomName"`
}

func (r *Room) Create() error{
	db := GetSelfDB()
	r.CreateAt = time.Now()
	return db.Create(&r).Error
}

func (r *Room) Update() error{
	db := GetSelfDB()
	return db.Model(r).Updates(r).Error
}


func DeleteRoom(id int64) error {
	db := GetSelfDB()
	e := Room{}
	e.ID = id
	return db.Delete(&e).Error
}

func GetRoom(id int64)(*Room,error) {
	db := GetSelfDB()
	u := &Room{}
	d := db.Where("id = ?",id).First(&u)
	return u,d.Error
}
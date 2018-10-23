package models

import (
	"time"
)

type Record struct {
	Model
	RoomName string `json:"roomname" gorm:"column:roomname"`   //会议室
	UserName string	 `json:"username" gorm:"column:username"`	//预订人
	Start time.Time `json:"start"` //开始时间
	End time.Time `json:"end"`      //结束时间
	Theme string `json:"theme" gorm:"column:theme"`//会议主题
	Member string `json:"member"` //参会人员
}

func (r *Record) Create() error{
	db := GetSelfDB()
	r.CreateAt = time.Now()
	return db.Create(&r).Error
}

func DeleteMeetInfo(id int64,u string) error{
	db := GetSelfDB()
	r := Record{}
	r.ID = id
	return db.Where(r.UserName).Delete(&r).Error
}

func (r *Record) Update() error{
	db := GetSelfDB()
	return db.Model(r).Updates(r).Error
}
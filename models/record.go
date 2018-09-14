package models

import (
	"time"
)

type Record struct {
	Model
	RoomRefer int64 `json:"rid" gorm:"column:rid"`   //会议室
	UserRefer int64	 `json:"uid" gorm:"column:uid"`	//预订人
	Start time.Time `json:"start"` //开始时间
	End time.Time `json:"end"`      //结束时间
	Participate string `json:"participate_num" gorm:"column:participate_num"`//可容纳人数
}
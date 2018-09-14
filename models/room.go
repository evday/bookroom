package models

type Room struct {
	Model
	Name string `json:"roomname" gorm:"column:roomname;not null"`  //会议室名称
	Location string `json:"location" gorm:"column:location;not null"`//位置
	Capacity string `json:"capacity" gorm:"column:capacity"` //可容纳人数
	State bool `json:"state" gorm:"column:state"`			//开放状态
	MeetRoom []Record `gorm:"ForeignKey:RoomRefer"`
}
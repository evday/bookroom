package models

type EquiRoom struct {
	Model 
	RoomId int64 `json:"roomid"`
	EquiId int64 `json:"Equimentid"`
	EquiNum int64 `json:"EquiNum"`
}
package models

import (
	"time"
	validator "gopkg.in/go-playground/validator.v9"
)

type Equipment struct {
	Model
	Name string `json:"equipname" gorm:"column:equipname;not null" binding:"required`  //设备名称
	Brand string `json:"brand" gorm:"column:brand"`					//设备品牌
	PModel string `json:"pmodel" gorm:"column:pmodel"`				//设备型号			//缩略图
	Num int64 `json:"num" gorm:"column:number" binding:"required`						//数量
}

//增加设备
func (e *Equipment) Create() error {
	db := GetSelfDB()
	e.CreateAt = time.Now()
	return db.Create(&e).Error
}

func (e *Equipment) FindAll() error {
	db := GetSelfDB()
	return db.Find(&e).Error
}

func DeleteEquip(id int64) error {
	db := GetSelfDB()
	e := Equipment{}
	e.ID = id
	return db.Delete(&e).Error

}

func (e *Equipment) Update() error {
	db := GetSelfDB()
	return db.Model(e).Updates(e).Error
}


func (e *Equipment) Validate() error {
	validate := validator.New()
	err := validate.Struct(e)
	return err
}


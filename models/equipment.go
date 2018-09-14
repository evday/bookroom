package models

type Equipment struct {
	Model
	Name string `json:"equipname" gorm:"column:equipname;not null"`  //设备名称
	Brand string `json:"brand" gorm:"column:brand"`					//设备品牌
	PModel string `json:"pmodel" gorm:"column:pmodel"`				//设备型号
	Image string `json:"image" gorm:"column:image"`					//缩略图
	Num int64 `json:"num" gorm:"column:number"`						//数量
}
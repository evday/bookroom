package models

import (

	"golang.org/x/crypto/bcrypt"
)


type User struct {
	Model
	Name string  `json:"username" gorm:"column:name;not null;unique"`   //姓名
	Password string	`json:"password" gorm:"column:password;not null"`	//密码
	IsAdmin bool	`json:"isadmin"`									//是否是管理员
	Records []Record `gorm:"ForeignKey:UserRefer"`
}

//保存之前给密码加密
func (u *User) BeforeSave() (err error) {
	var hash []byte
	hash, err = bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	u.Password = string(hash)
	return
}

//创建用户
func (u *User) Create() error {
	db := GetDB()
	return db.Create(&u).Error
}

//更新用户
func (u *User) Update() error {
	db := GetDB()
	return db.Model(u).Updates(u).Error
}

//删除用户
func DeleteUser(id uint64) error {
	db := GetDB()
	user := User{}
	user.ID = id
	return db.Delete(&user).Error
}

//获取用户
func GetUser(username string)(*User,error) {
	db := GetDB()
	u := &User{}
	d := db.Where("name = ?",username).First(&u)
	return u,d.Error
}

//获取用户列表
func GetListUser() ([]User,error) {
	db := GetDB()
	var users []User
	d := db.Find(&users)
	return users,d.Error
}
//验证密码
func (u *User) Compare(pwd string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pwd))
	return
}



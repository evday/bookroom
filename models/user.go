package models

import (
	"sync"
	"fmt"
	"time"
	"bookroom/pkg/constvar"
	"golang.org/x/crypto/bcrypt"
	validator "gopkg.in/go-playground/validator.v9"

)


type User struct {
	Model
	Name string  `json:"username" gorm:"column:name;not null;unique" binding:"required"`   //姓名
	Password string	`json:"password" gorm:"column:password;not null" binding:"required"`	//密码
	IsAdmin bool	`json:"isadmin"`									//是否是管理员
	Records []Record `gorm:"ForeignKey:UserName"`
}

type UserList struct {
	Lock *sync.Mutex
	IdMap map[int64]*User
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
	db := GetSelfDB()
	u.CreateAt = time.Now()
	return db.Create(&u).Error
}

//更新用户
func (u *User) Update() error {
	db := GetSelfDB()
	return db.Model(u).Updates(u).Error
}

//删除用户
func DeleteUser(id int64) error {
	db := GetSelfDB()
	user := User{}
	user.ID = id
	return db.Delete(&user).Error
}

//获取用户
func GetUser(username string)(*User,error) {
	db := GetSelfDB()
	u := &User{}
	d := db.Where("name = ?",username).First(&u)
	return u,d.Error
}

//获取用户列表
func ListUser(username string,offset,limit int)([]*User,int64,error)  {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}
	offset = (offset - 1)*limit
	users := make([]*User,0)

	var count int64
	fmt.Printf("username=%s offset=%d limit=%d \n",username,offset,limit)
	where := fmt.Sprintf("name like '%%%s%%'",username)
	if err := DB.Self.Model(&User{}).Where(where).Count(&count).Error;err != nil {
		return users,count,err
	}

	if err := DB.Self.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&users).Error;err != nil {
		return users,count,err
	}

	return users,count,nil
}
//验证密码
func (u *User) Compare(pwd string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pwd))
	return
}

func (u *User) Validate() error {
	validate := validator.New()
	err := validate.Struct(u)
	return err
}



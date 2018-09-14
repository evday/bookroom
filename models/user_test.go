package models

import (
	"testing"
)

func init() {
	SetDB()
}


func TestInsertUser(t *testing.T) {
	var user User
	user.Name = "evday"
	user.Password = "admin123"
	user.IsAdmin = true

	err := user.Create()
	if err != nil {
		t.Errorf("insert user failed,err:%v",err)
		return
	}
	t.Logf("insert user succ")
}

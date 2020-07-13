package mysql

import (
	"fmt"
	"testing"
)

type User struct {
	ID     int64  `json:"id" `
	Name   string `json:"name"`
	Cnname string `json:"cnname"`
	Passwd string `json:"-"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	IM     string `json:"im" gorm:"column:im"`
	QQ     string `json:"qq" gorm:"column:qq"`
	Role   int    `json:"role"`
}

func init(){
	InitDB()
}

func TestCreaterUser(t *testing.T){
	user := User{
		ID:     1,
		Name:   "hdf",
		Cnname: "何登峰",
		Passwd: "123456",
		Email:  "894220128@qq.com",
		Phone:  "15767634546",
		IM:     "test",
		QQ:     "894220128",
		Role:   0,
	}
	dt := Uic.Table("user").Create(&user)
	if dt.Error != nil {
		fmt.Println(dt.Error)
	}
}

func TestQueryUser(t *testing.T){
	//var user User
	//Uic.Table("user").Where("name = ?", "hdf").Scan(&user)
	user := User{Name: "hdf"}
	Uic.Table("user").Where(user).Scan(&user)
	if user.ID != 0 {
		fmt.Printf("%+v\n", user)
		return
	}
}

func TestUpdateUser(t *testing.T){
	//var user User
	//Uic.Table("user").Where("name = ?", "hdf").Scan(&user)
	uuser := User{Name: "hdf", Cnname: "何登丰"}
	dt := Uic.Table("user").Where("id = ?",1).Update(&uuser)
	if dt.Error != nil {
		fmt.Println(dt.Error)
		return
	}
	fmt.Println(uuser)
}

func TestDeleteUser(t *testing.T){
	//var user User
	//Uic.Table("user").Where("name = ?", "hdf").Scan(&user)
	uuser := User{}
	dt := Uic.Table("user").Where("id = ?",1).Delete(&uuser)
	if dt.Error != nil {
		fmt.Println(dt.Error)
		return
	}
	fmt.Println(uuser)
}
package mysql

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var  Uic     *gorm.DB
var  url = "root:123456@tcp(127.0.0.1:3306)/uic?charset=utf8&parseTime=True&loc=Local"


func InitDB(){
	var u *sql.DB

	var err error
	Uic, err = gorm.Open("mysql", url)
	if err != nil {
		fmt.Printf("connect to uic: %s", err.Error())
		return
	}
	Uic.Dialect().SetDB(u)
	Uic.LogMode(true)

	Uic.SingularTable(true)
}


func TestConnect(t *testing.T){
	InitDB()
}
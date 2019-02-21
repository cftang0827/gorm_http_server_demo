package model

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Gousers ..., This is the User model...
type Gousers struct {
	Userid string
	Name   string
	Gender string
	Age    int64
}

// var db *gorm.DB
// var err error
var db, err = gorm.Open("mysql", "test2:test2@tcp(192.168.79.11:41063)/cv?charset=utf8&parseTime=True&loc=Local")

// Init ...
func init() {
	fmt.Println("Connect to mysql db2")
	if err != nil {
		fmt.Println("Connection to db failed")
		panic(err)
	}
}

//Create ...
func (u *Gousers) Create() map[string]interface{} {
	err = db.Create(u).Error
	if err != nil {
		fmt.Println(err)
	}

	return map[string]interface{}{"userid": u.Userid, "name": u.Name, "gender": u.Gender, "age": u.Age, "errmesg": err}
}

//Read ...
func (u *Gousers) Read() map[string]interface{} {
	mesg := db.Where(map[string]interface{}{"userid": u.Userid}).Find(u)
	return map[string]interface{}{"userid": u.Userid, "name": u.Name, "gender": u.Gender, "age": u.Age, "errmesg": mesg.Error}

}

//Update ...
func (u *Gousers) Update() map[string]interface{} {
	updateinfo := map[string]interface{}{"userid": u.Userid, "name": u.Name, "gender": u.Gender, "age": u.Age}
	mesg := db.Model(u).Where(map[string]interface{}{"userid": u.Userid}).Updates(updateinfo)
	updateinfo["errmesg"] = mesg.Error
	return updateinfo
}

//Delete ....
func (u *Gousers) Delete() map[string]interface{} {
	mesg := db.Model(u).Where(map[string]interface{}{"userid": u.Userid}).Delete(u)
	deleteOutput := map[string]interface{}{"userid": u.Userid}
	deleteOutput["errmesg"] = mesg.Error
	return deleteOutput
}

package test

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"testing"
)

var db *gorm.DB

func TestDBConnect(t *testing.T) {

	fmt.Println("this is a functional testing !")
	var err error
	db, err = gorm.Open("mysql", "ad_service_rw:m03663LR04K1Ax5(10.16.39.52:3306)/ad_service?charset=utf8&parseTime=True&loc=Local")
	//db.DB().SetMaxIdleConns(5)
	//db.DB().SetMaxOpenConns(10)
	if err != nil {
		panic(err)
	}
}

func BenchmarkSimple(b *testing.B) {

	fmt.Println("this is a BenchMark test!")

}

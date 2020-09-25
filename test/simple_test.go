package test

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"testing"
	"time"
)

var db *gorm.DB

type Product struct {
	ID        int    `gorm:"primary_key"`
	Code      string `gorm:"type:varchar(20);"`
	Price     int    `gorm:"type:int;"`
	Name      string `gorm:"type:varchar(64);"`
	Mail      string `gorm:"type:varchar(256);"`
	CreatedAt time.Time
}

func TestDBConnect(t *testing.T) {

	fmt.Println("this is a functional testing !")
	var err error
	db, err = gorm.Open("mysql", "ads:Welcome2hush!@tcp(10.16.39.93:3306)/decontroler_dev?charset=utf8&parseTime=True&loc=Local")
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(10)
	if err != nil {
		panic(err)
	}

	//create table
	if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Product{}).Error; err != nil {
		panic(err)
	}

	//insert data
	//db.Create(&Product{Code: "10000", Price: 12, Name: "noodules"})

	//query data
	var product Product
	db.First(&product, 1)
	/*db.First(&product, "code = ?", "10000")

	//update data
	db.Model(&product).Update("Price", 2000)*/

	fmt.Println(product)
}

func BenchmarkSimple(b *testing.B) {

	fmt.Println("this is a BenchMark test!")

}

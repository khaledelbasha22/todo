package Sys

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"

)

func  PrintIfErr(err error) bool{
	if err != nil {
		log.Fatalf("Error is: %q\n", err)
		return  true
	}
	return false
}


func DB()  *gorm.DB{
	ConnectionString := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local","root","root","localhost","golang_tesk")
	conn, err := gorm.Open("mysql", ConnectionString)
	if err != nil {
		PrintIfErr(err)
		return nil
	}
	conn.DB().SetMaxOpenConns(900000) // Sane default
	conn.DB().SetMaxIdleConns(900000)
	conn.DB().SetConnMaxLifetime(time.Nanosecond)
	return conn
}


type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}





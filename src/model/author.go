package model

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name string `gorm:"varchar(50)"`
	Age  int
}

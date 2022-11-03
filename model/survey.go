package model

import (
	"github.com/jinzhu/gorm"
)

type Survey struct {
	gorm.Model
	Name string `json:"name"`
	Age  int64  `json:"age"`
	Hair string `json:"hair"`
}

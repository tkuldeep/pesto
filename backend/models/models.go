package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title string `json:"title" gorm:"text;not null;default:null"`
	Desc  string `json:"desc" gorm:"text;not null;default:null"`
}

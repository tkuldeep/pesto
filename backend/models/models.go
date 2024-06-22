package models

import "gorm.io/gorm"

// Task is model
type Task struct {
	gorm.Model
	Title      string `json:"title" gorm:"text;not null;default:null"`
	Desc       string `json:"desc" gorm:"text;default:null"`
	Status     int    `json:"-" gorm:"int; not null"`
	TaskStatus string `json:"status" gorm:"-"`
}

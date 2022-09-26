package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	EmailId  string `gorm:"unique;primaryKey"`
	Password string
	Resumes  []Resume
}

type Resume struct {
	gorm.Model
	Name   string
	UserID string
}

type Template struct {
	gorm.Model
	Name  string `gorm:"unique"`
	Url   string
	Tabs  string
	Struc string
}

type Usercreateac struct {
	Name     string `json:"name"`
	EmailId  string `json:"emailid"`
	Password string `json:"password"`
}

type Login struct {
	EmailId  string `json:"emailid"`
	Password string `json:"password"`
}

package model

import "gorm.io/gorm"

type Scholarchip struct {
	gorm.Model
	URL          string `gorm:"uniqueIndex"`
	Title        string
	Deadline     string
	Requirements string
	Country      string
	Language     string
}

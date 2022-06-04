package topic

import "gorm.io/gorm"

type Topic struct {
	gorm.Model
	Name         string `gorm:"unique"`
	ProfileImage *string
	Description  string
	Rules        *string
}

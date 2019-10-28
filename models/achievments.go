package models

import (
	"github.com/jinzhu/gorm"
)

type Achievement struct {
	gorm.Model
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Icon        string `json:"icon,omitempty"`
}

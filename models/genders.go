package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	u "sexsquare-api/utils"
)

type Gender struct {
	gorm.Model
	Name   string `json:"name,omitempty"`
}

func (gender *Gender) Validate() (map[string]interface{}, bool) {

	if gender.Name == "" {
		return u.Message(false, "Gender name should be on the payload"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (gender *Gender) Create() map[string]interface{} {

	if resp, ok := gender.Validate(); !ok {
		return resp
	}

	GetDB().Create(gender)

	resp := u.Message(true, "success")
	resp["gender"] = gender
	return resp
}

func GetGender(id uint) *Gender {

	gender := &Gender{}
	err := GetDB().Table("genders").Where("id = ?", id).First(gender).Error
	if err != nil {
		return nil
	}
	return gender
}

func GetGenders() []*Gender {

	genders := make([]*Gender, 0)
	err := GetDB().Table("genders").Find(&genders).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return genders
}

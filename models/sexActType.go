package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	u "sexsquare-api/utils"
)

type SexActType struct {
	gorm.Model
	Title   string `json:"title,omitempty"`
}

func (sexActType *SexActType) Validate() (map[string]interface{}, bool) {

	if sexActType.Title == "" {
		return u.Message(false, "Sex act type title should be on the payload"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (sexActType *SexActType) Create() map[string]interface{} {

	if resp, ok := sexActType.Validate(); !ok {
		return resp
	}

	GetDB().Create(sexActType)

	resp := u.Message(true, "success")
	resp["sexActType"] = sexActType
	return resp
}

func GetSexActType(id uint) *SexActType {

	sexActType := &SexActType{}
	err := GetDB().Table("sexActTypes").Where("id = ?", id).First(sexActType).Error
	if err != nil {
		return nil
	}
	return sexActType
}

func GetSexActTypes() []*SexActType {

	sexActTypes := make([]*SexActType, 0)
	err := GetDB().Table("sexActTypes").Find(&sexActTypes).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return sexActTypes
}

package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	u "sexsquare-api/utils"
)

type SexAct struct {
	gorm.Model
	User           *Account      `json:"User,omitempty" gorm:"ForeignKey: Email"`
	Types          []*SexActType `json:"Types,omitempty"`
	Contraceptives []string      `json:"Contraceptives,omitempty"`
	Note           string        `json:"Note,omitempty"`
	Partners       []*Partner    `json:"Partners,omitempty"`
	Place          *Place        `json:"Place,omitempty"`
	StartSex       string        `json:"StartSex,omitempty"`
	FinishSex      string        `json:"FinishSex,omitempty"`
	Private        bool          `json:"Private,omitempty"`
}

func (sexAct *SexAct) Validate() (map[string]interface{}, bool) {

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (sexAct *SexAct) Create() map[string]interface{} {

	if resp, ok := sexAct.Validate(); !ok {
		return resp
	}

	GetDB().Create(sexAct)

	resp := u.Message(true, "success")
	resp["sexAct"] = sexAct
	return resp
}

func GetSexAct(id uint) *SexAct {

	sexAct := &SexAct{}
	err := GetDB().Table("sexActs").Where("id = ?", id).First(sexAct).Error
	if err != nil {
		return nil
	}
	return sexAct
}

func GetSexActs(id uint) []*SexAct {

	sexActs := make([]*SexAct, 0)
	err := GetDB().Table("sexActs").Where("id = ?", id).Find(&sexActs).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return sexActs
}

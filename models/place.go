package models

import (
	"github.com/jinzhu/gorm"
	u "sexsquare-api/utils"
)

type Place struct {
	gorm.Model
	Address   string  `json:"Address,omitempty"`
	Host      string  `json:"Host,omitempty"`
	LatLng    *LatLng `json:"LatLng,omitempty"`
}

func (place *Place) Validate() (map[string]interface{}, bool) {

	if place.Address == "" {
		return u.Message(false, "Address should be on the payload"), false
	}

	if place.Host == "" {
		return u.Message(false, "Host should be on the payload"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

package models

type Partner struct {
	PartnerId uint    `json:"PartnerId,omitempty" gorm:"ForeignKey ID"`
	StartSex string   `json:"StartSex,omitempty"`
}



package model

type Qipp struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Txt    string `json:"t"`
	UserId string `json:"-"`
}

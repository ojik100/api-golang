package models

type Dosen struct {
	Id   uint   `gorm:"primaryKey"`
	Nidn string `json:"nidn"`
	Nama string `json:"nama"`
}

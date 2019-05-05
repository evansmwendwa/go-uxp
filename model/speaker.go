package model

// Speaker - speaker model
type Speaker struct {
	Base
	Name     string `json:"name"`
	Position string `json:"position"`
	Company  string `json:"company"`
	Bio      string `gorm:"size:1500" json:"bio"`
	LinkedIn string `gorm:"type:varchar(100);size:255" json:"linkedin"`
	Twitter  string `gorm:"size:255" json:"twitter"`
}

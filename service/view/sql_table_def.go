package view

//
type BaseId struct {
	Id int `gorm:"primary_key" json:"-"`
}

package models

type User struct {
	Basic
	FirstName string  `json:"firstName" gorm:"type:varchar(255)"`
	LastName  string  `json:"lastName" gorm:"type:varchar(255)"`
	Email     *string `json:"email" gorm:"uniqueIndex, type:varchar(255)"`
	Password  []byte  `json:"-"`
}

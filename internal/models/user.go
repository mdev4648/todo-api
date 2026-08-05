package models

import "gorm.io/gorm"

type User struct {
	gorm.Model //Actually it's embedding another struct. it contains fields like ID, CreatedAt, UpdatedAt, and DeletedAt. By embedding gorm.Model, the User struct inherits these fields and their associated behavior. ID uint , CreateAt time.Time, UpdatedAt time.Time, DeletedAt gorm.DeletedAt

	Name     string `gorm:"size:100;not null"` //This is a struct tag that provides metadata about the Name field. It specifies that the maximum size of the Name field is 100 characters and that it cannot be null in the database.
	Email    string `gorm:"size:100;unique;not null"`
	Password string `gorm:"not null"`
}

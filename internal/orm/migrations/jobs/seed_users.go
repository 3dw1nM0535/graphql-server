package jobs

import (
	"github.com/3dw1nM0535/go-gql-server/internal/orm/models"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

var (
	uname                    = "John Doe"
	fname                    = "John"
	lname                    = "Doe"
	nname                    = "johndoe"
	description              = "I'm cool"
	location                 = "21, Maragoli"
	firstUser   *models.User = &models.User{
		Email:       "example@email.com",
		Name:        &uname,
		FirstName:   &fname,
		LastName:    &lname,
		NickName:    &nname,
		Description: &description,
		Location:    &location,
	}
)

// SeedUsers inserts a dummy user
var SeedUsers *gormigrate.Migration = &gormigrate.Migration{
	ID: "SEED_USERS",
	Migrate: func(db *gorm.DB) error {
		return db.Create(&firstUser).Error
	},
	Rollback: func(db *gorm.DB) error {
		return db.Delete(&firstUser).Error
	},
}

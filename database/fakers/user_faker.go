package fakers

import (
	"time"

	"github.com/arif-rizal1122/go-online-shop/app/models"
	"github.com/bxcodec/faker/v4"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// go get -u github.com/bxcodec/faker/v4
// go get -u github.com/gosimple/slug



func UserFaker(db *gorm.DB) *models.User  {
	return &models.User{
		ID: 			uuid.New().String(),
		FirsName:       faker.FirstName(),
		LastName: 		faker.LastName(),
		Email:			faker.Email(),
		Password: 		"",
		RememberToken:  "",
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},					
		DeletedAt: 		gorm.DeletedAt{},
	}
}
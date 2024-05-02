package fakers

import (
	"log"
	"time"

	"github.com/arif-rizal1122/go-online-shop/app/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func ProductImageFakers(db *gorm.DB) *models.ProductImage {
     productID := ProductFaker(db)
	 err := db.Create(&productID).Error
	 if err != nil {
		log.Fatal(err)
	 }
	 
	 return &models.ProductImage{
		ID:             uuid.New().String(),
		ProductID:      productID.ID,
		Path:           "img/products/",
		ExtraLarge:     "",
		Large:          "",
		Medium:         "",
		Small:          "",
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
	 }

	 
}
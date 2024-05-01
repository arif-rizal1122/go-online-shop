package seeders

import (
	"fmt"

	"github.com/arif-rizal1122/go-online-shop/database/fakers"
	"gorm.io/gorm"
)



type Seeder struct {
	Seeder interface{}
}



func RegisterSeeder(db *gorm.DB) []Seeder {
   return []Seeder{
       {Seeder: fakers.UserFaker(db)},
	   {Seeder: fakers.ProductFaker(db)},
   }

}



func DBSeed(db *gorm.DB) error {
	for _, seeder := range RegisterSeeder(db) {
		err := db.Debug().Create(seeder.Seeder).Error
        if err != nil {
			return err
		}
	}
	fmt.Println("db seeder successfully")
	return nil
}
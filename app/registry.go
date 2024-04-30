package app

import "github.com/arif-rizal1122/go-online-shop/app/models"




type Model struct {
	Model interface{}
}



func RegisterModels() []Model {
	return []Model{
		{Model: models.User{}},
		{Model: models.Address{}},		
	}
}

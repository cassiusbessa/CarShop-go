package mocks

import (
	entity "github.com/cassiusbessa/goCarShop/internal/db/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Car = entity.ICar{
	Model:    "Gol",
	Year:     2010,
	Color:    "Black",
	Status:   true,
	BuyValue: 20000.00,
	DoorsQty: 4,
	SeatsQty: 5,
}

var ErrorCar = entity.ICar{
	Model:    "Gol",
	Year:     2010,
	Color:    "Black",
	Status:   true,
	BuyValue: 20000.00,
	DoorsQty: 0,
	SeatsQty: 0,
}

var Seeds = []entity.ICar{
	{
		Id:       primitive.NewObjectID(),
		Model:    "Gol",
		Year:     2010,
		Color:    "Blue",
		Status:   true,
		BuyValue: 20000.00,
		DoorsQty: 4,
		SeatsQty: 5,
	},
	{
		Id:       primitive.NewObjectID(),
		Model:    "Astra",
		Year:     2015,
		Color:    "Silver",
		Status:   true,
		BuyValue: 30000.00,
		DoorsQty: 4,
		SeatsQty: 5,
	},
	{
		Id:       primitive.NewObjectID(),
		Model:    "Celta",
		Year:     2012,
		Color:    "Black",
		Status:   true,
		BuyValue: 15000.00,
		DoorsQty: 4,
		SeatsQty: 5,
	},
}

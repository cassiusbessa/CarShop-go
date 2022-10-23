package mocks

import entity "github.com/cassiusbessa/goCarShop/internal/db/entities"

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

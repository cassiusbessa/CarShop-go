package entity

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ICar struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Model    string             `json:"model,omitempty" bson:"model,omitempty" validate:"required"`
	Year     int                `json:"year,omitempty" bson:"year,omitempty" validate:"required"`
	Color    string             `json:"color,omitempty" bson:"color,omitempty" validate:"required"`
	Status   bool               `json:"status,omitempty" bson:"status,omitempty" validate:"required"`
	BuyValue float64            `json:"buyValue,omitempty" bson:"buyValue,omitempty" validate:"required"`
	DoorsQty int                `json:"doorsQty,omitempty" bson:"doorsQty,omitempty" validate:"required,gte=1,lte=5"`
	SeatsQty int                `json:"seatsQty,omitempty" bson:"seatsQty,omitempty" validate:"required,gte=1,lte=10"`
}

func (c *ICar) Validate() error {
	valid := validator.New()
	err := valid.Struct(c)
	if err != nil {
		return err
	}
	return nil
}

func (c *ICar) UpdateValidate(fields ...string) error {
	valid := validator.New()
	err := valid.StructPartial(c, fields...)
	if err != nil {
		return err
	}
	return nil
}

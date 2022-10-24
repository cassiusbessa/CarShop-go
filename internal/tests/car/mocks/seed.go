package mocks

import (
	context "context"
)

func Seeder() {
	for _, car := range Seeds {
		TestCollection.InsertOne(context.Background(), car)
	}
}

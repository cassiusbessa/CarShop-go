package mocks

import (
	context "context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newCollection(db, collection string) *mongo.Collection {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	return client.Database(db).Collection(collection)
}

var TestCollection = newCollection("test", "cars")

package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCollection *mongo.Collection

func Connect(uri, db, collection string) {
	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	if err = mongoClient.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	}
	MongoCollection = mongoClient.Database(db).Collection(collection)
}

func Insert(i any) {
	if _, err := MongoCollection.InsertOne(context.TODO(), i); err != nil {
		log.Fatal(err)
	}
}


// input value is empty string means query by field exists
func Query[T any](key, value string) []T {
	var results []T
	var filter bson.D
	if value != "" {
		filter = bson.D{{Key: key, Value: value}}
	} else {
		filter = bson.D{{Key: key, Value: bson.D{{Key: "$exists", Value: true}}}}
	}
	cur, _ := MongoCollection.Find(context.TODO(), filter)
	if err := cur.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	return results
}

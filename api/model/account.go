package model

import (
	"api/db"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Account struct {
	Id   string `bson:"id" json:"id"`
	Name string `bson:"_id" json:"name"`
}

func GetAccountsByMessages() (*[]bson.M, error) {
	collection := db.GetMongoCollection(nil)
	pipeline := mongo.Pipeline{
		bson.D{
			{"$group", bson.D{
				{"_id", "$owner_name"},
				{"owner_name", bson.D{
					{"$first", "$owner_name"},
				}},
				{"owner_id", bson.D{
					{"$first", "$owner_id"},
				}},
				{"count", bson.D{
					{"$sum", 1},
				}},
			}},
		},
		bson.D{
			{"$sort", bson.D{{"count", -1}}},
		},
		bson.D{
			{"$limit", 10},
		},
	}

	cursor, err := collection.Aggregate(nil, pipeline)

	var results []bson.M

	if err = cursor.All(nil, &results); err != nil {
		log.Fatal(err)
	}

	return &results, err
}

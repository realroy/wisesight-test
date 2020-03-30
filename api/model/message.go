package model

import (
	"api/db"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Message struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}

func GetMessageCount() (int64, error) {
	collection := db.GetMongoCollection(nil)

	return collection.CountDocuments(nil, bson.D{})
}

func GetMessageByEngagements() (*[]bson.M, error) {
	collection := db.GetMongoCollection(nil)
	pipeline := mongo.Pipeline{
		bson.D{{
			"$project", bson.M{
				"message":    1,
				"engagement": 1,
			},
		}},
		bson.D{{"$sort", bson.M{"engagement": -1}}},
		bson.D{{"$limit", 10}},
	}

	cursor, err := collection.Aggregate(nil, pipeline)

	var results []bson.M

	if err = cursor.All(nil, &results); err != nil {
		log.Fatal(err)
	}

	return &results, err
}

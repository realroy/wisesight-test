package model

import (
	"api/db"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Word struct {
	Key   string `json:"text"`
	Total uint64 `json:"value"`
}

func GetWords() (interface{}, error) {
	collection := db.GetMongoCollection(nil)
	pipeline := mongo.Pipeline{
		bson.D{{"$project", bson.D{{"words", 1}}}},
		bson.D{{"$unwind", "$words"}},
		bson.D{
			{"$group", bson.D{
				{"_id", "$words"},
				{"text", bson.D{{"$first", "$words"}}},
				{"value", bson.D{{"$sum", 1}}},
			},
			},
		},
		bson.D{{"$sort", bson.D{{"value", -1}}}},
		bson.D{{"$limit", 100}},
	}

	cursor, err := collection.Aggregate(nil, pipeline)

	var results []bson.M

	if err = cursor.All(nil, &results); err != nil {
		log.Fatal(err)
	}

	return &results, err
}

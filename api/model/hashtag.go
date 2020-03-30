package model

import (
	"api/db"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Hashtag struct {
	Key   string `json:"text"`
	Total uint64 `json:"value"`
}

func GetHashtags() (interface{}, error) {
	collection := db.GetMongoCollection(nil)
	pipeline := mongo.Pipeline{
		bson.D{{"$project", bson.D{{"hashtags", 1}}}},
		bson.D{{"$unwind", "$hashtags"}},
		bson.D{
			{"$group", bson.D{
				{"_id", "$hashtags"},
				{"text", bson.D{{"$first", "$hashtags"}}},
				{"value", bson.D{
					{"$sum", 1},
				},
				},
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

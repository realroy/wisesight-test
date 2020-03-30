package config

import "os"

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		return defaultValue
	}

	return value
}

func Port() string {
	return ":" + getEnv("PORT", "8080")
}

func MongoURI() string {
	return getEnv("MONGO_URI", "mongodb://localhost:27017")
}

func MongoDBName() string {
	return getEnv("MONGO_DB_NAME", "tkdev19")
}

func MongoCollectionName() string {
	return getEnv("MONGO_COLLECTION_NAME", "wisesightTest")
}

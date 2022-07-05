package database

import "go.mongodb.org/mongo-driver/mongo"

func DBInstance() *mongo.Client {
	err := godotenv.Load()
}

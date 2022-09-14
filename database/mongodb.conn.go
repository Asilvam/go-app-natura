package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

func GetCollection(collection string) *mongo.Collection {

	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}
	uriMongodb := os.Getenv("URI_MONGODB")
	database := os.Getenv("DATABASE")

	client, err := mongo.NewClient(options.Client().ApplyURI(uriMongodb))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client.Database(database).Collection(collection)
}

package database

import (
	"context"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func MongoDBCursor() (*mongo.Client, error) {
	return mongo.Connect(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
}

func MongoInsertOne(collection_name string, document interface{}) (*mongo.InsertOneResult, error) {
	client, err := MongoDBCursor()
	if err != nil {
		return nil, err
	}

	collection := client.Database("mydb").Collection(collection_name)

	defer client.Disconnect(context.Background())
	return collection.InsertOne(context.Background(), document)
}

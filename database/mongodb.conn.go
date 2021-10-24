package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

//remember in change this variable when start the project

var (
	usr 		= "ygor"
	pwd 		= "123456"
	host 		= "localhost"
	port 		= 27017
	database 	= "testforwex"
)

func GetCollection(collection string) *mongo.Collection {

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", usr, pwd, host, port)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		panic(err.Error())
	}

	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	err = client.Connect(ctx)

	if err != nil {
		panic(err.Error())
	}

	return client.Database(database).Collection(collection)
}

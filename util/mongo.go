package util

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var MongoClient *mongo.Client

func init() {
	var err error
	MongoClient, err = mongo.Connect(
		options.
			Client().
			ApplyURI("mongodb+srv://szh:shizihao1996@cluster0.u7opsq0.mongodb.net/?appName=cluster0").
			SetServerAPIOptions(options.ServerAPI(options.ServerAPIVersion1)),
	)
	if err != nil {
		panic(err)
	}
}

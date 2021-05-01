package gzmongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GZMongo struct {
	address  string
	password string

	ctx    context.Context
	Client *mongo.Client
}

func New(address, password string) *GZMongo {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://bd:%v@%v", password, address)))

	if err != nil {
		log.Fatal("can not connect to mongoDB due to error: ", err)
	}

	return &GZMongo{
		address:  address,
		password: password,

		Client: client,
		ctx:    ctx,
	}
}

// disconnect mongo db
func (gm *GZMongo) Disconnect() {
	if err := gm.Client.Disconnect(gm.ctx); err != nil {
		panic(err)
	}
}

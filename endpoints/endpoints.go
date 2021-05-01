package endpoints

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// endpoint db name
var DB_NAME string = "devdb"
var COLLECTION_NAME string = "endpoints"

// struct for inventory
type Endpoint struct {
	client *mongo.Client
	ctx    context.Context
}

// create new inventory instance
func New(client *mongo.Client, ctx context.Context) *Endpoint {
	return &Endpoint{
		client: client,
		ctx:    ctx,
	}
}

func (e *Endpoint) GetEndpoints() ([]*EndpoointList, error) {
	database := e.client.Database(DB_NAME)

	cln := database.Collection(COLLECTION_NAME)

	var eps []*EndpoointList

	// cursor, err := cln.Find(i.ctx, bson.D{})
	cursor, err := cln.Find(e.ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	err = cursor.All(e.ctx, &eps)
	if err != nil {
		return nil, err
	}

	return eps, err
}

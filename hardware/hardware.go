package hardware

import (
	"context"

	"github.com/javadmohebbi/gravityzone-extra/endpoints"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// database name
var DB_NAME string = "deviceControlEvents"

// collection name
var COLLECTION_NAME string = "detectedDevices"

// struct for hardware
type Hardware struct {
	client *mongo.Client
	ctx    context.Context
}

// create new hardware instance
func New(client *mongo.Client, ctx context.Context) *Hardware {
	return &Hardware{
		client: client,
		ctx:    ctx,
	}
}

// get all detected hardware list
func (h *Hardware) GetAll() ([]*HardwareList, error) {
	var hwds []*HardwareList

	hwds, err := h.getHwds()
	if err != nil {
		return nil, err
	}

	e := endpoints.New(h.client, h.ctx)
	epsList, _ := e.GetEndpoints()

	for _, hw := range hwds {

		for _, ep := range epsList {
			for _, e := range hw.EndpointIDs {
				if ep.ID == e {
					hw.Endpoints = append(hw.Endpoints, *ep)
				}
			}
		}
	}

	return hwds, nil

}

// get all hardwares in the database
func (h *Hardware) getHwds() ([]*HardwareList, error) {
	database := h.client.Database(DB_NAME)

	cln := database.Collection(COLLECTION_NAME)

	var hwdList []*HardwareList

	cursor, err := cln.Find(h.ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	err = cursor.All(h.ctx, &hwdList)
	if err != nil {
		return nil, err
	}

	return hwdList, err
}

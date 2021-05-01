package inventory

import (
	"context"
	"strings"

	"github.com/javadmohebbi/gravityzone-extra/endpoints"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// database name
var DB_NAME string = "applicationControl"

// collection name
var COLLECTION_NAME string = "applicationInventory"

// collection name - endpoint mapping
var COLLECTION_NAME_MAPPING string = "applicationInventoryEndpointMappings"

// endpoint db name
var DB_NAME_ENDPOINT string = "devdb"
var COLLECTION_NAME_ENDPOINT string = "endpoints"

// struct for inventory
type Inventory struct {
	client *mongo.Client
	ctx    context.Context
}

// create new inventory instance
func New(client *mongo.Client, ctx context.Context) *Inventory {
	return &Inventory{
		client: client,
		ctx:    ctx,
	}
}

// return all application in inventory
func (i *Inventory) GetAll() ([]*InventoryList, error) {
	var apps []*InventoryList
	// var inventoryList []InventoryList

	apps, err := i.getApps()
	if err != nil {
		return nil, err
	}

	e := endpoints.New(i.client, i.ctx)
	epsList, _ := e.GetEndpoints()

	for _, ap := range apps {
		ap.PathString = i.extractPath(ap.Path)
		epsByApp, err := i.getEndpintsByApp(ap.ID, epsList)
		if err == nil {
			ap.Endpoints = epsByApp
		}
	}

	return apps, err
	// return inventoryList, err

}

// get all apps in the inventory
func (i *Inventory) getApps() ([]*InventoryList, error) {
	database := i.client.Database(DB_NAME)

	cln := database.Collection(COLLECTION_NAME)

	var ivl []*InventoryList

	// cursor, err := cln.Find(i.ctx, bson.D{})
	cursor, err := cln.Find(i.ctx, bson.M{
		"specifics.hash": bson.M{
			"$ne": nil,
		},
	})

	if err != nil {
		return nil, err
	}

	err = cursor.All(i.ctx, &ivl)
	if err != nil {
		return nil, err
	}

	return ivl, err
}

// get endpoint using this app
func (i *Inventory) getEndpintsByApp(objId primitive.ObjectID, epsList []*endpoints.EndpoointList) ([]*EndpointMapping, error) {
	database := i.client.Database(DB_NAME)

	cln := database.Collection(COLLECTION_NAME_MAPPING)

	var eps []*EndpointMapping

	// cursor, err := cln.Find(i.ctx, bson.D{})
	cursor, err := cln.Find(i.ctx, bson.M{
		"nodeId": objId,
	})

	if err != nil {
		return nil, err
	}

	err = cursor.All(i.ctx, &eps)
	if err != nil {
		return nil, err
	}

	for _, e := range eps {
		for _, ee := range epsList {
			if e.EndpointID == ee.ID {
				e.EndpointName = ee.Name
				e.EndpointOS = ee.OperatingSystemVersion
			}
		}
	}

	return eps, err
}

// extract path from / delimited path of mongo db
func (i *Inventory) extractPath(path string) string {
	database := i.client.Database(DB_NAME)

	cln := database.Collection(COLLECTION_NAME)

	paths := strings.Split(path, "/")

	var retStr = ""

	for _, s := range paths {
		var il InventoryList

		docID, err := primitive.ObjectIDFromHex(s)
		if err != nil {
			continue
		}

		// cursor, err := cln.Find(i.ctx, bson.D{})
		res := cln.FindOne(i.ctx, bson.M{"_id": docID})

		if err != nil {
			continue
		}

		err = res.Decode(&il)
		if err != nil {
			continue
		}

		retStr += il.Name + "/"

	}

	return retStr
}

// func (i *Inventory) getChildren(objId primitive.ObjectID) ([]InventoryList, error) {
// 	database := i.client.Database(DB_NAME)

// 	cln := database.Collection(COLLECTION_NAME)

// 	var ivl []InventoryList

// 	// cursor, err := cln.Find(i.ctx, bson.D{})
// 	cursor, err := cln.Find(i.ctx, bson.M{
// 		// "$and": []bson.M{
// 		// {"parentId": objId},
// 		// {"specifics": bson.M{"$ne": nil}},
// 		"parentId": objId},
// 	)

// 	if err != nil {
// 		return nil, err
// 	}

// 	err = cursor.All(i.ctx, &ivl)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return ivl, err
// }

// func (i *Inventory) getParents() ([]*InventoryList, error) {
// 	database := i.client.Database(DB_NAME)

// 	cln := database.Collection(COLLECTION_NAME)

// 	var ivl []*InventoryList

// 	// cursor, err := cln.Find(i.ctx, bson.D{})
// 	cursor, err := cln.Find(i.ctx, bson.M{
// 		"$or": []bson.M{
// 			{"parentId": nil},
// 			{"specifics": nil},
// 		},
// 	})

// 	if err != nil {
// 		return nil, err
// 	}

// 	err = cursor.All(i.ctx, &ivl)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return ivl, err
// }

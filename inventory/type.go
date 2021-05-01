package inventory

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InventoryList struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`

	ParentID primitive.ObjectID `bson:"parentId,omitempty"`

	Name string `bson:"name,omitempty"`

	Path string `bson:"path,omitempty"`

	PathString string `bson:"-"`

	Specifics InventorySpecifics `bson:"specifics,omitempty"`

	Children []InventoryList `bson:"-"`

	// got from endpoints
	// EndpointOS   string
	// EndpointName string
	Endpoints []*EndpointMapping `bson:"-"`
}

type EndpointMapping struct {
	NodeID     primitive.ObjectID `bson:"nodeId,omitempty"`
	EndpointID primitive.ObjectID `bson:"endpointId,omitempty"`

	EndpointOS   string `bson:"-"`
	EndpointName string `bson:"-"`
}

type InventorySpecifics struct {
	Hash         string    `bson:"hash,omitempty"`
	DiscoveredOn time.Time `bson:"discoveredOn,omitempty"`
	Version      string    `bson:"version,omitempty"`
}

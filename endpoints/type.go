package endpoints

import "go.mongodb.org/mongo-driver/bson/primitive"

type EndpoointList struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`

	Name string `bson:"name,omitempty"`

	OperatingSystemVersion string `bson:"operatingSystemVersion,omitempty"`
}

package hardware

import (
	"time"

	"github.com/javadmohebbi/gravityzone-extra/endpoints"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HardwareList struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`

	DeviceID string `bson:"deviceId,ommitempty"`

	DeviceName string `bson:"deviceName,ommitempty"`

	EndpointIDs []primitive.ObjectID `bson:"endpointIds,ommitempty"`

	LastDetectionDate time.Time `bson:"lastDetectionDate,omitempty"`

	Endpoints []endpoints.EndpoointList `bson:"-"`
}

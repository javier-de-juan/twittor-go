package responseModel

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Tweet struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID    string             `bson:"user_id" json:"user_id,omitempty"`
	Text      string             `bson:"text" json:"text,omitempty"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at,omitempty"`
}

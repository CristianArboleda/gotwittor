package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// FollowersTweetsResponse : Object to response for to get all followers tweets
type FollowersTweetsResponse struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID         string             `bson:"userid" json:"userid,omitempty"`
	RelationUserID string             `bson:"relationuserid" json:"relationuserid,omitempty"`
	Tweet          struct {
		ID         string    `bson:"_id" json:"_id,omitempty"`
		Message    string    `bson:"message" json:"message,omitempty"`
		CreateDate time.Time `bson:"createdate" json:"createdate,omitempty"`
	}
}

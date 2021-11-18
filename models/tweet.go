package models

import "time"

type Tweet struct {
	ID         string    `bson:"_id" json:"_id,omitempty"`
	UserID     string    `bson:"userid" json:"userid,omitempty"`
	Message    string    `bson:"message" json:"message,omitempty"`
	CreateDate time.Time `bson:"createdate" json:"createdate,omitempty"`
}

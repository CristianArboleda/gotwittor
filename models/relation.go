package models

import "time"

type Relation struct {
	UserID         string    `bson:"userid" json:"userid,omitempty"`
	RelationUserID string    `bson:"relationuserid" json:"relationuserid,omitempty"`
	CreateDate     time.Time `bson:"createdate" json:"createdate,omitempty"`
}

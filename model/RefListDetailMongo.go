package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ReferenceListDetailMongo struct {
	ID          int64     `json:"_id" bson:"_id"`
	Name        string    `json:"name"`
	Code        string    `json:"code"`
	ReferenceListId int64 `json:"referenceListId"`
	ReferenceListName string `json:"referenceListName"`
	DateCreated time.Time `json:"dateCreated"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	Slug primitive.Binary `json:"slug"`
}
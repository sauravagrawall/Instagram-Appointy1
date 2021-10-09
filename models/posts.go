package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Posts struct{
	Id			bson.ObjectId 	`json:"id" bson:"_id"`
	Caption		string			`json:"caption" bson:"caption"`
	ImageURL	string			`json:"imgurl" bson:"imgurl"`
	Time		string			`json:"time" bson:"time"`
}
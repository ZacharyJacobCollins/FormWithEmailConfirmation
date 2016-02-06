package models

import "gopkg.in/mgo.v2/bson"

type (
	// User represents the structure of our resource
	User struct {
		Id     	   bson.ObjectId `json:"id" bson:"_id"`
		Name 			 string     	 `json:"name" bson:"name"`
		Username   string        `json:"username" bson:"username"`
		Email  		 string        `json:"email" bson:"email"`
	}
)

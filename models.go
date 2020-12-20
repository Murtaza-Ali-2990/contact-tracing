package models
package timestamp

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"github.com/path/to/timestamp"
)

//Create Struct
type User struct {
	ID     		primitive.ObjectID 		`json:"_id,omitempty" bson:"_id,omitempty"`
	Name   		string             		`json:"name,omitempty" bson:"name,omitempty"`
	DOB  		time.Date           	`json:"dob" bson:"dob,omitempty"`
	Phone 		string            		`json:"phone" bson:"phone,omitempty"`
	Email		string					`json:"email" bson:"email,omitempty"`
	Created_at  *timestamp.Timestamp  	`bson:"created_at,omitempty" json:"created_at,omitempty"`
}

type Contact struct {
	ID1 		string 					`json:"firstname,omitempty" bson:"firstname,omitempty"`
	ID2  		string 					`json:"lastname,omitempty" bson:"lastname,omitempty"`
	Created_at  timestamp.Timestamp  	`bson:"created_at,omitempty" json:"created_at,omitempty"`
}
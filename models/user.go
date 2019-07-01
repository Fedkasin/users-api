package models

import (
	"time"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"firstname" bson:"firstname"`
	LastName  string             `json:"lastname" bson:"lastname"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
	CreatedAt time.Time 		 `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time 		 `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

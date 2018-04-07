package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	//User struct represents users of the application
	User struct {
		ID           bson.ObjectId `bson:"_id,omitempty" json:"id"`
		FirstName    string        `json:"firstname"`
		LastName     string        `json:"lastname"`
		Email        string        `json:"email"`
		Password     string        `json:"password,omitempty"`
		HashPassword []byte        `json:"hashpassword,omitempty`
	}
	//Task struct used to represent task added by user
	Task struct {
		ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		CreatedBy   string        `json:"createdby"`
		Name        string        `json:"name"`
		Description string        `json:"description"`
		CreatedOn   time.Time     `json:"createdon,omitempty"`
		Due         time.Time     `json:"due,omitempty"`
		Status      string        `json:"status,omitempty"`
		Tags        []string      `json:"tags,omitempty"`
	}
	//TaskNote struct used to represent notesd added by user
	TaskNote struct {
		ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		TaskID      bson.ObjectId `bson:"taskid,omitempty" json:"taskid"`
		Description string        `json:"description"`
		CreatedOn   time.Time     `json:"createdon,omitempty"`
	}
)

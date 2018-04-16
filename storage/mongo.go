package storage

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
)

const DBName string = "talk"
const MongoURL string = "localhost"

// NewSession returns a new mongoDB session
func NewSession() (*mgo.Session, error) {
	session, err := mgo.Dial(MongoURL)
	if err != nil {
		return nil, fmt.Errorf("error on MongoDB connection: %q", err)
	}

	return session, nil
}

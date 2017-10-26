package main

import (
	"fmt"
	"os"

	mgo "gopkg.in/mgo.v2"
)

func Connect() (session *mgo.Session) {
	connectURL := "url"
	session, err := mgo.Dial(connectURL)
	if err != nil {
		fmt.Printf("Cannot connect to MongoDB, %v\n", err) os.Exit(1)
	}

	session.SetSafe(&mgo.Safe{})

	return session
}

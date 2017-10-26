package main

import (
	"fmt"
	"os"
	"time"

	mgo "gopkg.in/mgo.v2"
)

func Connect() (session *mgo.Session) {

	var (
		MongoDBHosts = os.Getenv("MONGO_URI")
		AuthDatabase = "fp-journal"
		AuthUserName = os.Getenv("MONGO_USER")
		AuthPassword = os.Getenv("MONGO_PASS")
	)

	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{MongoDBHosts},
		Timeout:  600 * time.Second,
		Database: AuthDatabase,
		Username: AuthUserName,
		Password: AuthPassword,
	}

	// Connect to DB with credentials, temporarily using
	// mLab for this initial implementation
	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		fmt.Printf("Cannot connect to MongoDB, %v\n", err)
		os.Exit(1)
	}

	session.SetSafe(&mgo.Safe{})

	return session
}

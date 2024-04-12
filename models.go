package main

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	client *mongo.Client
)

type User struct {
	S_ID    string `json:"s_id" bson:"s_id"`
	IDNO    string `json:"idno" bson:"idno"`
	SNAME   string `json:"sname" bson:"sname"`
	HOSCODE string `json:"hoscode" bson:"hoscode"`
	ROOM    string `json:"room" bson:"room"`
	// BMAIL   string    `json:"bmail" bson:"bmail"`
}

func init() {
	// Initialize MongoDB client
	if c, err := StartDB(); err != nil {
		panic(err)
	} else {
		client = c
	}
}

func InsertUser(user User) error {
	// Get the users collection
	collection := client.Database("GrubDatabase").Collection("users")

	// Insert the user into the collection
	_, err := collection.InsertOne(context.Background(), user)
	return err
}

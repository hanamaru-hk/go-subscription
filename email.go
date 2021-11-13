package main

import "go.mongodb.org/mongo-driver/bson/primitive"

// Task - Model of a basic task
type Email struct {
	ID    primitive.ObjectID
	Email string
}

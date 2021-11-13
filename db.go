package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// Timeout operations after N seconds
	connectTimeout           = 5
	connectionStringTemplate = "mongodb+srv://%s:%s@%s/hanamaru?retryWrites=true&w=majority"

	// connectionStringTemplate = "mongodb://%s:%s@%s"
)

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

// GetConnection Retrieves a client to the MongoDB
func getConnection() (*mongo.Client, context.Context, context.CancelFunc) {
	username := goDotEnvVariable("MONGODB_USERNAME")
	password := goDotEnvVariable("MONGODB_PASSWORD")
	clusterEndpoint := goDotEnvVariable("MONGODB_ENDPOINT")

	log.Printf("username: %s", username)
	log.Printf("password: %s", password)
	log.Printf("clusterEndpoint: %s", clusterEndpoint)

	connectionURI := fmt.Sprintf(connectionStringTemplate, username, password, clusterEndpoint)

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))
	if err != nil {
		log.Printf("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Printf("Failed to connect to cluster: %v", err)
	}

	// Force a connection to verify our connection string
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("Failed to ping cluster: %v", err)
	}

	fmt.Println("Connected to MongoDB!")
	return client, ctx, cancel
}

// GetAll Retrives all emails from the db
func GetAll() ([]*Email, error) {
	var emails []*Email

	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	db := client.Database("hanamaru")
	collection := db.Collection("subscription")
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &emails)
	if err != nil {
		log.Printf("Failed marshalling %v", err)
		return nil, err
	}
	return emails, nil
}

//Create creating a email in a mongo
func Create(email *Email) (primitive.ObjectID, error) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	email.ID = primitive.NewObjectID()

	result, err := client.Database("hanamaru").Collection("subscription").InsertOne(ctx, email)
	if err != nil {
		log.Printf("Could not create email: %v", err)
		return primitive.NilObjectID, err
	}
	oid := result.InsertedID.(primitive.ObjectID)
	return oid, nil
}

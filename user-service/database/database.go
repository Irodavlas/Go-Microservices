package database

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/irodavlas/user-service/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type IDatabase interface {
	PingDatabase() error
	Create(User *model.User) error
	Delete(username string) error
	Update(User *model.User) error
	Get(user model.User) (*model.User, error)
}
type Database struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewDatabaseConnection(uri, dbName, collectionName string) (*Database, error) {
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	collection := client.Database(collectionName).Collection(collectionName)
	log.Println("[INFO] Created database connection")
	return &Database{
		client:     client,
		collection: collection,
	}, nil

}
func (db *Database) PingDatabase() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := db.client.Ping(ctx, nil)
	if err != nil {
		return err
	}
	log.Println("[INFO] Seccesfully Pinged Database")
	return nil
}
func (d *Database) Create(user *model.User) error {
	_, err := d.collection.InsertOne(
		context.TODO(),
		user)
	if err != nil {
		log.Printf("[ERROR]:%s", err.Error())
		return err
	}
	log.Printf("[INFO] Create operation succesfull")
	return nil
}

// check if err is not nil and also the user.
func (d *Database) Get(user model.User) (*model.User, error) {
	var dbUser model.User
	filter := bson.M{"username": user.Username}

	err := d.collection.FindOne(context.Background(), filter).Decode(&dbUser)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Printf("[INFO] No user found with username: %s", user.Username)
			return nil, nil // Not found, return nil to indicate absence
		}
		log.Printf("[ERROR] Failed to query user: %v", err)
		return nil, fmt.Errorf("database query error: %w", err)
	}

	log.Printf("[INFO] User '%s' found", dbUser.Username)
	return &dbUser, nil
}
func (d *Database) Delete(username string) error {
	filter := bson.M{"username": username}
	result, err := d.collection.DeleteOne(
		context.TODO(),
		filter,
	)
	if err != nil {
		log.Printf("[ERROR]:%s", err.Error())
		return err
	}
	if result.DeletedCount == 0 {
		log.Printf("[INFO] Number of deleted rows:%d", result.DeletedCount)
		return fmt.Errorf("documents deleted:%d", result.DeletedCount)
	}
	log.Printf("[INFO] Documents deleted:%d", result.DeletedCount)
	return nil
}

// this needs to expand as we must be able to change password or username
func (d *Database) Update(user model.User) error {

	result, err := d.collection.UpdateOne(
		context.TODO(),
		user.Username,
		user,
	)
	if err != nil {
		log.Println("[ERROR]:", err.Error())
		return err
	}
	if result.MatchedCount == 0 {
		log.Println("[ERROR]: No documentes found to be updated")
		return fmt.Errorf("document updated:%d", result.MatchedCount)
	}
	return nil
}

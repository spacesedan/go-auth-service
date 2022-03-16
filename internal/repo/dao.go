package repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
)

type DAO interface {
	NewUserQuery() UserQuery
}

type dao struct {
}

var DB *mongo.Database

// NewDAO create a data access object
func NewDAO(db *mongo.Database) DAO {
	DB = db
	return &dao{}
}

// NewMongoDB connect to mongoose
func NewMongoDB() (*mongo.Database, error) {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalln("Error loading .env file")
			return nil, err
		}
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatalln("Error reading uri: ", err)
	}

	ctx := context.Background()

	err = client.Connect(ctx)
	if err != nil {
		return nil, errors.New("failed to connect to mongo")
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("connected to mongo")

	db := client.Database(os.Getenv("DB"))

	return db, nil
}

// NewUserQuery interact with the user repository
func (d *dao) NewUserQuery() UserQuery {
	return &userQuery{}
}

package repo

import (
	"authentication/internal/models"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	collectionName = "users"
	ctx            = context.Background()
)

type UserQuery interface {
	FindOneByEmail(email string) (*models.User, error)
	InsertNewUser(email, passwordHash string) (*models.User, error)
}

type userQuery struct {
}

// FindOneByEmail find a user using an email
func (u *userQuery) FindOneByEmail(email string) (*models.User, error) {
	m := DB.Collection(collectionName)

	filter := bson.M{
		"email": email,
	}

	var user *models.User

	result := m.FindOne(ctx, filter)
	err := result.Decode(&user)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("no user found")
	}

	return user, nil
}

func (u *userQuery) InsertNewUser(email, passwordHash string) (*models.User, error) {
	m := DB.Collection(collectionName)

	insert := bson.M{
		"$set": bson.M{
			"email":    email,
			"password": passwordHash,
		},
	}

	_, err := m.InsertOne(ctx, insert)
	if err != nil {
		return nil, errors.New("failed to insert user")
	}

	var user *models.User

	filter := bson.M{
		"email": email,
	}

	result := m.FindOne(ctx, filter)
	if result == nil {
		return nil, errors.New("no user found")
	}

	err = result.Decode(&user)
	if err != nil {
		return nil, errors.New("failed to decode user")
	}

	return user, nil
}

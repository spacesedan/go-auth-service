package repo

import (
	"authentication/internal/apperrors"
	"authentication/internal/models"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	collectionName = "users"
	ctx            = context.Background()
)

type UserQuery interface {
	FindOneByEmail(email string) (*models.User, error)
	InsertNewUser(email, passwordHash string) error
}

type userQuery struct {
}

// FindOneByEmail find a user using an email
func (u *userQuery) FindOneByEmail(email string) (*models.User, error) {
	m := DB.Collection(collectionName)

	fmt.Println("FIND", email)
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
		return nil, apperrors.NoUserErr
	}

	return user, nil
}

func (u *userQuery) InsertNewUser(email, passwordHash string) error {
	m := DB.Collection(collectionName)

	insert := bson.M{
		"email":     email,
		"password":  passwordHash,
		"username":  "",
		"lastName":  "",
		"firstName": "",
	}

	_, err := m.InsertOne(ctx, insert)
	if err != nil {
		return errors.New("failed to insert user")
	}

	return nil

}

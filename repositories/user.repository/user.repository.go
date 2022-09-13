package user_repository

import (
	"context"
	"github.com/Asilvam/go-app-natura.git/database"
	"github.com/Asilvam/go-app-natura.git/models"
	"go.mongodb.org/mongo-driver/bson"
)

var collection = database.GetCollection("users")
var ctx = context.Background()

func CreateUser(user models.User) error {
	return nil
}

func GetUsers() (models.Users, error) {
	var users models.Users
	filter := bson.D{}

	cur, err := collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {

		var user models.User
		err = cur.Decode(&user)

		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil

}

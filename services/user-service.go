package services

import (
	"context"
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"

	"github.com/Fedkasin/users-api/config"
	"github.com/Fedkasin/users-api/models"
	"github.com/Fedkasin/users-api/mongodb"
)

const CollectionName = "users"

var conf = config.Config

func GetUserById(id primitive.ObjectID, user *models.User) (error) {
	collection := mongodb.Client.Database(conf.DbName).Collection(CollectionName)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err:= collection.FindOne(ctx, bson.D{{"_id", id}}).Decode(&user)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func UpdateUserById(id primitive.ObjectID, user models.User) mongodb.SingleResult {
	collection := mongodb.Client.Database(conf.DbName).Collection(CollectionName)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	updatedUser := collection.FindOneAndUpdate(ctx, bson.D{{"_id", id}}, bson.D{
		{ "$set", bson.D{
			{ "firstname", user.FirstName },
			{"lastname", user.LastName},
			{ "email", user.Email },
			{ "password", user.Password }},
		},
		{"$currentDate", bson.D{{"updated_at", true}}} }, nil)

	return updatedUser
}

func RemoveUserById() ([]models.User, error) {
	return nil, nil
}

func CreateUser(user models.User) (mongodb.InsertOneResult, error) {
	collection := mongodb.Client.Database(conf.DbName).Collection(CollectionName)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return result, nil
}

func GetUsers(users []models.User) ([]models.User, error) {
	collection := mongodb.Client.Database(conf.DbName).Collection(CollectionName)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user models.User
		cursor.Decode(&user)
		users = append(users, user)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return users, nil
}
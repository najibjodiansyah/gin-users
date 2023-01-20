package services

import (
	"context"
	"errors"
	"log"

	"github.com/najibjodiansyah/gin-users-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(usercollection *mongo.Collection, ctx context.Context) *UserServiceImpl {
	return &UserServiceImpl{
		usercollection: usercollection,
		ctx:            ctx,
	}
}

func (u *UserServiceImpl) Get() ([]*models.User, error) {
	var users []*models.User
	cursor, err := u.usercollection.Find(u.ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(u.ctx)
	for cursor.Next(u.ctx) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, mongo.ErrNoDocuments
	}
	return users, nil
}

func (u *UserServiceImpl) GetByUser(name string) (*models.User, error) {
	var user *models.User
	query := bson.D{bson.E{Key: "name", Value: name}}
	err := u.usercollection.FindOne(u.ctx, query).Decode(&user)
	return user, err
}

func (u *UserServiceImpl) Create(user *models.User) error {
	_, err := u.usercollection.InsertOne(u.ctx, user)
	return err
}

func (u *UserServiceImpl) Update(name string, user *models.User) error {
	filter := bson.D{bson.E{Key: "name", Value: name}}
	updateQuery := bson.D{bson.E{
		Key: "$set",
		Value: bson.D{
			bson.E{Key: "name", Value: user.Name},
			bson.E{Key: "email", Value: user.Email},
			bson.E{Key: "age", Value: user.Age},
			bson.E{Key: "address", Value: user.Address},
			bson.E{Key: "phone", Value: user.Phone}},
	}}
	log.Println(updateQuery)
	update := bson.D{bson.E{Key: "$set", Value: user}}
	result, err := u.usercollection.UpdateOne(u.ctx, filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount != 1 {
		return errors.New("user not found")
	}
	return nil
}

func (u *UserServiceImpl) Delete(name string) error {
	filter := bson.D{bson.E{Key: "name", Value: name}}
	result, err := u.usercollection.DeleteOne(u.ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount != 1 {
		return errors.New("user not found")
	}
	return nil
}

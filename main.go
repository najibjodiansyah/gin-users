package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/najibjodiansyah/gin-users-api/controllers"
	"github.com/najibjodiansyah/gin-users-api/services"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server         *gin.Engine
	userService    services.UserService
	userController controllers.UserController
	ctx            context.Context
	userCollection *mongo.Collection
	mongoClient    *mongo.Client
	err            error
)

func init() {
	ctx = context.TODO()
	mongoConn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoClient, err = mongo.Connect(ctx, mongoConn)
	if err != nil {
		log.Fatal(err)
	}
	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")
	userCollection = mongoClient.Database("userdb").Collection("users")
	userService = services.NewUserService(userCollection, ctx)
	userController = controllers.NewUserController(userService)
	server = gin.Default()
}

func main() {
	defer mongoClient.Disconnect(ctx)
	path := server.Group("/api/v1")
	userController.RegisterRoutes(path)
	log.Fatal(server.Run(":8080"))
}

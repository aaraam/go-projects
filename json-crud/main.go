package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"json-crud/controllers"
)

func main() {
	r := httprouter.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client, err := getSession()
	if err != nil {
		log.Fatal(err)
	}
	uc := controllers.NewUserController(client)
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	r.PUT("/user/:id", uc.UpdateUser)

	fmt.Println("Server started at port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

func getSession() (*mongo.Client, error) {
	uri := os.Getenv("MONGO_URI")
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB")
	return client, nil
}

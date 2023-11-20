package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func register(c echo.Context) error {

	item := new(User)
	if err := c.Bind(item); err != nil {
		return err

	}
	// users = append(users, *item)

	collection := DB.Database("UserService").Collection("Users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	password, err := bcrypt.GenerateFromPassword([]byte(item.Password), 14)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)

	}
	res, err := collection.InsertOne(ctx, bson.D{{"username", item.Username}, {"password", password}})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)

	}
	id := res.InsertedID

	resp := fmt.Sprint("User registered successfully %v", id)
	return c.JSON(http.StatusCreated, resp)

}

func login(c echo.Context) error {
	item := new(User)
	if err := c.Bind(item); err != nil {
		return err
	}
	collection := DB.Database("UserService").Collection("Users")

	filter := bson.D{{"username", item.Username}}

	result := collection.FindOne(context.TODO(), filter)

	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			fmt.Println("record does not exist")
			return c.JSON(http.StatusNotFound, map[string]string{"error": "record does not exist"})
		}
		panic(result.Err())
	}

	var user User
	if err := result.Decode(&user); err != nil {
		panic(err)
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(item.Password))
	if err == nil {
		resp := fmt.Sprintf("successfully loggedin %v", user.Username)
		return c.JSON(http.StatusOK, resp)

	} else {
		return c.JSON(http.StatusOK, "incorrect username or password")

	}

}

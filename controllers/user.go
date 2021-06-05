package controllers

import (
	"context"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	l        *log.Logger
	database *mongo.Database
	ctx      context.Context
}

type Book struct {
	Name        string
	PublisherID string
	Cost        string
	StartTime   string
	EndTime     string
}

func (u *User) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		rw.Write([]byte("Simple get request"))
	}

	if r.Method == http.MethodPost {
		testCollection := u.database.Collection("BooksRead")

		inserRes, err := testCollection.InsertOne(context.TODO(), Book{Name: "Some Potter", PublisherID: "IBN123", Cost: "1232", StartTime: "2013-10-01T01:11:18.965Z", EndTime: "2013-10-01T01:11:18.965Z"})
		log.Println("InsertResponse : ", inserRes)
		log.Println("Error : ", err)

	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func NewUser(l *log.Logger, database *mongo.Database, ctx context.Context) *User {
	return &User{l, database, ctx}
}

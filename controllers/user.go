package controllers

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/buger/jsonparser"
	"go.mongodb.org/mongo-driver/bson"
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
		paths := [][]string{
			{"uuid"},
			{"tz"},
			{"ua"},
			{"st"},
		}

		result := parseJSON(r.Body, paths)

		testCollection := u.database.Collection("BooksRead")

		inserRes, err := testCollection.InsertOne(context.TODO(), bson.D{
			{"ua", result.ua},
			{"tz", result.tz},
			{"uuid", result.uuid},
		})

		log.Println("InsertResponse : ", inserRes)
		log.Println("Error : ", err)

	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

type SmallPayload struct {
	tz   string
	st   string
	ua   string
	uuid string
}

var data SmallPayload

func parseJSON(body io.ReadCloser, paths [][]string) SmallPayload {
	b, _ := ioutil.ReadAll(body)
	jsonparser.EachKey(b, func(idx int, value []byte, vt jsonparser.ValueType, err error) {
		switch idx {
		case 0:
			data.uuid = string(value)
		case 1:
			v, _ := jsonparser.ParseString(value)
			data.tz = string(v)
		case 2:
			data.ua = string(value)
		case 3:
			v, _ := jsonparser.ParseString(value)
			data.st = string(v)
		}
	}, paths...)

	return data
}

func NewUser(l *log.Logger, database *mongo.Database, ctx context.Context) *User {
	return &User{l, database, ctx}
}

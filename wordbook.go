package main

import (
    "log"
    "context"
    "time"
    
    "github.com/hoonhoons/hoon-vocabot-go/formats"
    
    "net/http"
    "github.com/labstack/echo"
    //"github.com/labstack/echo/middleware"
    
    //"github.com/mongodb/mongo-go-driver/bson"
    //"github.com/mongodb/mongo-go-driver/bson/primitive"
    
    "github.com/mongodb/mongo-go-driver/bson"
    "github.com/mongodb/mongo-go-driver/mongo"
)

func putWordToWordbook(c echo.Context) error {
    request := new(formats.Request)
    c.Bind(&request)
    userId := request.UserRequest.User.Id
    wordId := lastWord[userId].ID // mongo
    
    // Connect
    ctx, _ := context.WithTimeout(context.Background(), 30 * time.Second)
    client, err := mongo.Connect(ctx, "mongodb://localhost:27017")

    if err != nil {
        log.Fatal(err)
    }
    
    // Check the connection
    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    
    db := client.Database("wordbot")
    wordbook := db.Collection("wordbook")
    _, err = wordbook.InsertOne(ctx, bson.M{"userID": userId, "wordID": wordId})
    
    if err != nil {
        log.Fatal(err)
    }
    
    d := &formats.Data {
        Msg1: "완료",
    }
    
    r := &formats.Response {
        Version: "2.0",
        Data: d,
    }
    
    return c.JSON(http.StatusOK, r)
}
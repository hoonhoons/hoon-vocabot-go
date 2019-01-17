/*
Park, Seonghoon
*/

package main

import (
    //"fmt"
    "net/http"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    
    "github.com/hoonhoons/hoon-vocabot-go/formats"
    //"github.com/mongodb/mongo-go-driver/mongo"
)

var lastWord = map[string]formats.Word{}
var lastSentence = map[string]string{}


func main() {
    // Echo instance
    
    //fmt.Println("Hoon-vocabot has started!")
    
    e := echo.New()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Init
    //initQuiz()
    
    // Routes
    e.GET("/", hello)
    e.POST("/quiz", generateQuiz)
    e.GET("/quiz", sendQuiz)
    e.POST("/check", checkQuiz)
    e.POST("/put", putWordToWordbook)

    // Start server
    e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
    return c.String(http.StatusOK, "Hello, World!")
}



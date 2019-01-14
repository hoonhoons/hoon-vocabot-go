/*
Park, Seonghoon
*/

package main

import (
    //"fmt"
    "net/http"
    "github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	
	//"github.com/mongodb/mongo-go-driver/mongo"
)

type Response struct {
    Version string `json:"version"`
    Template *Template `json:"template,omitempty"`
    Data *Data `json:"data,omitempty"`
}

type Template struct {
    Outputs []Component `json:"outputs"`
    //QuickReplies []QuickReply `json:"quicReplies,omitempty"`
}

type Component struct {
    SimpleText *SimpleText `json:"simpleText,omitempty"`
}

type SimpleText struct {
    Text string `json:"text"`
}

type Data struct {
    Msg1 string `json:"msg1,omitempty"`
    Msg2 string `json:"msg2,omitempty"`
    Msg3 string `json:"msg3,omitempty"`
    Msg4 string `json:"msg4,omitempty"`
    Msg5 string `json:"msg5,omitempty"`
}


func main() {
	// Echo instance
	
	//fmt.Println("Hoon-vocabot has started!")
	
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.POST("/quiz", sendQuiz)
	e.GET("/quiz", generateQuiz)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}



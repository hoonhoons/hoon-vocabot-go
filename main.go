/*
Park, Seonghoon
*/

package main

import (
    //"fmt"
    "net/http"
    
    "github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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
    Body string `json:"body,omitempty"`
    A string `json:"a,omitempty"`
    B string `json:"b,omitempty"`
    C string `json:"c,omitempty"`
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
	e.POST("/quiz", quiz)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}


func quiz(c echo.Context) error {
    /*
    d := &Data {
        Number: "1",
        TotalNumber: "1",
        Body: "test?",
        A: "A!",
        B: "B!",
        C: "C!",
    }
    */
    
    /*
    t := &Template {
        Outputs: []Component{
            Component {
                &SimpleText {
                    Text: "test!",
                },
            },
        },
    }
    
    r := &Response {
        Version: "2.0",
        Template: t,
    }
    */
    
    d := &Data {
        Body: "test?",
        A: "A!",
        B: "B!",
        C: "C!",
    }
    
    r := &Response {
        Version: "2.0",
        Data: d,
    }
    
    return c.JSON(http.StatusOK, r)
}
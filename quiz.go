/*
Park, Seonghoon
*/

package main

import (
    "log"
    "context"
    //"fmt"
    "time"
    
    "net/http"
    "github.com/labstack/echo"
	//"github.com/labstack/echo/middleware"
	
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type Quiz struct {
    Body string
    A string
    B string
    C string
    D string
}

type Word struct {
    Word string
    Pos string
    Meaning string
    Examples []string
}

func sendQuiz(c echo.Context) error {
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
        Msg1: "test?",
        Msg2: "A!",
        Msg3: "B!",
        Msg4: "C!",
    }
    
    r := &Response {
        Version: "2.0",
        Data: d,
    }
    
    return c.JSON(http.StatusOK, r)
}

func generateQuiz(c echo.Context) error {
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
    teps := db.Collection("teps")

    // Randomly select a word (단, pos가 비어있거나, 예문이 없는 경우는 제외)
    // 감탄사 등 고려해야 (TODO)
    var answer Word
    for {
        pipeline := mongo.Pipeline{
            {{"$sample", bson.D{{"size", 1}}}},
        }
        
        cur, err := teps.Aggregate(ctx, pipeline)
        if err != nil {
        	log.Fatal(err)
        }
        
        defer cur.Close(ctx)
        for cur.Next(ctx) {
            err := cur.Decode(&answer)
            if err != nil {
            	log.Fatal(err)
            }
        }
        
        // 적당한 예제 없으면 넘겨야됨 (TODO)
        if len(answer.Examples) == 0 {
            continue
        }
        
        if (answer.Pos == "") {
            continue
        }
        
        break
    }
    
    // 정답과 같은 품사의 단어 1개 찾기 (4개 TODO)
    // 중복 방지 (TODO)

    var others [4]Word
    for i := 0; i < 4; i++ {
        for {
            pipeline := mongo.Pipeline{
                {{"$match", bson.D{{"pos", answer.Pos}}}},
                {{"$sample", bson.D{{"size", 1}}}},
            }
            
            cur, err := teps.Aggregate(ctx, pipeline)
            if err != nil {
            	log.Fatal(err)
            }
            
            defer cur.Close(ctx)
            for cur.Next(ctx) {
                err := cur.Decode(&others[i])
                if err != nil {
                	log.Fatal(err)
                }
            }
            
            // 적당한 예제 없으면 넘겨야됨 (TODO)
            if len(others[i].Examples) == 0 {
                continue
            }
            
            if (others[i].Pos == "") {
                continue
            }
            
            break
        }
    }


    // ***********************************************************

    /*
    //filter := bson.M{"word": "vitiate"}
    //err = teps.FindOne(ctx, filter).Decode(&result)
    fmt.Println(result.Word)
    fmt.Println(result.Pos)
    fmt.Println(result.Examples[1])
    */
    
    /*
    cur, err := teps.Find(ctx, nil)
    if err != nil {
        fmt.Println("si...bal...")
    	log.Fatal(err)
    }
    
    defer cur.Close(ctx)
    for cur.Next(ctx) {
        var result bson.M
        err := cur.Decode(&result)
        if err != nil {
        	log.Fatal(err)
        }
    	fmt.Println("hello")
    }
    
    if err := cur.Err(); err != nil {
        log.Fatal(err)
    }
    */
    
    // ***********************************************************
    
    
    // Disconnect
    err = client.Disconnect(ctx)

    if err != nil {
        log.Fatal(err)
    }
    
    resultString := answer.Word + " " + answer.Pos
    for i := range others {
        resultString += "\n" + others[i].Word + " " + others[i].Pos
    }
    
    return c.String(http.StatusOK, resultString)
}
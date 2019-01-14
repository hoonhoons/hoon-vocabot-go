/*
Park, Seonghoon
*/

package main

import (
    "log"
    "context"
    "time"
    "strings"
    "math/rand"
    "fmt"
    
    "github.com/hoonhoons/hoon-vocabot-go/formats"
    
    "net/http"
    "github.com/labstack/echo"
	//"github.com/labstack/echo/middleware"
	
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

/*
type Quiz struct {
    Body string
    A string
    B string
    C string
    D string
}
*/


// 스킬 서버 -> 봇 서버
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

type Word struct {
    Word string
    Pos string
    Meaning string
    Examples []string
}

var lastWord = map[string]Word{}
var lastSentence = map[string]string{}

func initQuiz() {
    // 봇 실행시 초기화되어야 하는 부분
    //lastQuiz := map[string]Word{} // 여기서 초기화하면 안되더라
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
    request := new(formats.Request)
	c.Bind(&request)
	userId := request.UserRequest.User.Id
	
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
    var question string
    
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
        
        // 검색 실패한 단어이므로 넘겨야됨 (빈 내용)
        if (answer.Pos == "") {
            continue
        }
        
        // 적당한 예제 없으면 넘겨야됨 (TODO)
        if len(answer.Examples) == 0 {
            continue
        }
        
        // 예문 랜덤 셔플
        rand.Seed(time.Now().UnixNano())
        rand.Shuffle(len(answer.Examples), func(i, j int) {
            answer.Examples[i], answer.Examples[j] = answer.Examples[j], answer.Examples[i]
        })
        
        // 빈칸 뚫기 (es, ed 처리 TODO)
        for _, example := range(answer.Examples) {
            if !strings.Contains(example, answer.Word) {
                continue
            }
            
            question = strings.Replace(example, answer.Word, "__________", -1)
            lastSentence[userId] = example
            break
        }
        
        if question == "" {
            continue
        }
        
        break
    }
    
    // 정답과 같은 품사의 단어 1개 찾기 (4개 TODO)
    // 중복 방지 (TODO)
    var others [5]Word
    for i := 0; i < 4; i++ {
        count := 0
        
        for {
            
            count += 1
            if count > 10 { // 무한루프 방지
                break
            }
            
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
            
            // 정답과 중복 테스트
            if (others[i].Word == answer.Word) {
                continue
            }
            
            // 다른 오답과의 중복 테스트 TODO
            
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
    
    
    // 정답 저장
    lastWord[userId] = answer
    fmt.Println(userId) // debug
    
    
    // 선택지 랜덤으로 섞기
    others[4] = answer
    rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(others), func(i, j int) {
        others[i], others[j] = others[j], others[i]
    })

    
    quizMsg := fmt.Sprintf("Q. %s\n\nㄱ) %s\nㄴ) %s\nㄷ) %s\nㄹ) %s\nㅁ) %s\n",
        question, others[0].Word, others[1].Word, others[2].Word, others[3].Word, others[4].Word,
    )
    
    d := &Data {
        Msg1: quizMsg,
    }
    
    r := &Response {
        Version: "2.0",
        Data: d,
    }
    
    return c.JSON(http.StatusOK, r)
}


func checkQuiz(c echo.Context) error {
    request := new(formats.Request)
	c.Bind(&request)
	userId := request.UserRequest.User.Id
	
	fmt.Println(userId) // debug
    
    answerMsg := lastSentence[userId] + "\n\n"
    answerMsg += "'" + lastWord[userId].Word + "'\n"
    answerMsg += lastWord[userId].Pos + "\n"
    answerMsg += lastWord[userId].Meaning
    
    d := &Data {
        Msg1: answerMsg,
    }
    
    r := &Response {
        Version: "2.0",
        Data: d,
    }
    
    return c.JSON(http.StatusOK, r)
}
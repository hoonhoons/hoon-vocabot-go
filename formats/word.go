package formats

import (
    "github.com/mongodb/mongo-go-driver/bson/primitive"
)

// 단어
type Word struct {
    ID primitive.ObjectID `bson:"_id"`
    Word string
    Pos string
    Meaning string
    Examples []string
}
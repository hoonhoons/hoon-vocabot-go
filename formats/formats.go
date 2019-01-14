package formats

type Request struct {
    UserRequest *UserRequest `json:"userRequest,omitempty"`
    Bot *Bot `json:"bot,omitempty"`
    Action *Action `json:"action,omitempty"`
}

type UserRequest struct {
    Timezone string `json:"timezone,omitempty"`
    Block *Block `json:"block,omitempty"`
    Utterance string `json:"utterance,omitempty"`
    Lang string `json:"lang,omitempty"`
    User *User `json:"user,omitempty"`
}

type Bot struct {
    Id string `json:"id,omitempty"`
    Name string `json:"name,omitempty"`
}

type Action struct {
    Id string `json:"id,omitempty"`
    Name string `json:"name,omitempty"`
    // TODO params, detailParams, clientExtra
}

type Block struct {
    Id string `json:"id,omitempty"`
    Name string `json:"name,omitempty"`
}

type User struct {
    Id string `json:"id,omitempty"` // 사용자 식별 ID
    Type string `json:"type,omitempty"`
    // properties TODO
}
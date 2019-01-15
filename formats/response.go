package formats

// 스킬 서버 --> 봇 서버
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
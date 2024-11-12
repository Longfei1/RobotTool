package jsmsg

type ShowMessage struct {
	Id        uint64      `json:"id"`        //唯一id
	ProtoId   uint16      `json:"protoId"`   //消息号
	ProtoType string      `json:"protoType"` //消息类型
	Msg       interface{} `json:"msg"`       //消息内容
	Type      UiClassType `json:"type"`
	Result    string      `json:"result"`
}

type UiClassType string

const (
	Primary UiClassType = "primary"
	Success UiClassType = "success"
	Info    UiClassType = "info"
	Warning UiClassType = "warning"
	Danger  UiClassType = "danger"
)

type TipTag struct {
	Name string      `json:"name"`
	Type UiClassType `json:"type"`
	Msg  string      `json:"msg"`
}

type ProtoMsg struct {
	Id           interface{} `json:"id"` //唯一标识
	Name         string      `json:"name"`
	Schema       interface{} `json:"schema"`       //描述schema的对象
	DefaultValue interface{} `json:"defaultValue"` //默认值
}

type RequestMsgData struct {
	Id       interface{} `json:"id"` //消息号，对应ProtoMsg.Id
	JsonData string      `json:"jsonData"`
}

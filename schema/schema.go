package schema

import (
	"encoding/json"
)

type Meta struct {
	WsId   string `json:"ws_id"`
	UserId int64  `json:"user_id"`
}

type Message struct {
	FromMicroservice string          `json:"from_microservice"`
	ToMicroservice   string          `json:"to_microservice"`
	RequestId        string          `json:"request_id"`
	ResponseId       string          `json:"response_id"`
	Method           string          `json:"method"`
	Version          string          `json:"version"`
	Status           Status          `json:"status"`
	Meta             Meta            `json:"meta"`
	Body             json.RawMessage `json:"body"`
}

type MessageHandler func(message Message)

func GenerateFullMethodNameFromMessage(message Message) string {
	return message.Method + "-" + message.Version
}

func FillMessageForSend(requestId string, fromMs string, toMs string, status Status, message *Message) {
	message.FromMicroservice = fromMs
	message.ToMicroservice = toMs
	message.ResponseId = requestId
	message.Status = status
}

func GenerateResponseMessage(requestMessage Message) Message {
	return Message{
		FromMicroservice: requestMessage.ToMicroservice,
		ToMicroservice:   requestMessage.FromMicroservice,
		ResponseId:       requestMessage.RequestId,
		Method:           requestMessage.Method,
		Version:          requestMessage.Version,
		Meta:             requestMessage.Meta,
	}
}

func FillMetaForMessage(message *Message, meta Meta) {
	message.Meta = meta
}

func FillStatusForMessage(message *Message, status Status) {
	message.Status = status
}

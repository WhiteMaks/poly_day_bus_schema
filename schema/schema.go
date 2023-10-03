package schema

import (
	"encoding/json"
)

const (
	PathMsGateway          = "gateway"
	PathMsAuthenticate     = "authenticate"
	PathMsDeveloperService = "developer-service"
)

const (
	Version10 = "1.0"
)

type Status int64

const (
	OK             Status = 2000
	BAD_REQUEST    Status = 4000
	AUTHORIZED     Status = 4010
	NOT_FOUND      Status = 4040
	INTERNAL_ERROR Status = 5000
)

const (
	PathSubscriptionDeveloperServiceAuthenticateUserCreated10 = PathMsDeveloperService + "." + PathEventAuthenticateUserCreated10
)

var Microservices = []string{
	PathMsGateway,
	PathMsAuthenticate,
	PathMsDeveloperService,
}

var Events = []string{
	PathEventAuthenticateUserCreated10,
}

var Subscribers = map[string][]string{
	PathEventAuthenticateUserCreated10: {
		PathSubscriptionDeveloperServiceAuthenticateUserCreated10,
	},
}

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

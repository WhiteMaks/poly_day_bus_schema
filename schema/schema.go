package schema

import "encoding/json"

const (
	PathMsGateway          = "gateway"
	PathMsAuthenticate     = "authenticate"
	PathMsDeveloperService = "developer-service"
)

const (
	PathEventAuthenticateUserCreated10 = PathMsAuthenticate + ".event.user-created-1.0"
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

type Message struct {
	FromMicroservice string          `json:"from_microservice"`
	ToMicroservice   string          `json:"to_microservice"`
	Method           string          `json:"method"`
	Version          string          `json:"version"`
	Body             json.RawMessage `json:"body"`
}

type MessageHandler func(message Message)

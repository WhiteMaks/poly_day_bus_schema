package schema

import "encoding/json"

const (
	PathMsGateway          = "gateway"
	PathMsAuthenticate     = "authenticate"
	PathMsDeveloperService = "developer-service"
)

const (
	Version10 = "1.0"
)

const (
	EventNameUserCreated = "user-created"
)

const (
	PathEventAuthenticateUserCreated10 = PathMsAuthenticate + ".event." + EventNameUserCreated + "-" + Version10
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

type UserCrated10 struct {
	ID   int64  `json:"id"`
	Role string `json:"role"`
}

func GenerateMessageForUserCreated10(id int64, role string) (Message, error) {
	body := UserCrated10{
		ID:   id,
		Role: role,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return Message{}, err
	}

	message := Message{
		Method:  EventNameUserCreated,
		Version: Version10,
		Body:    bodyBytes,
	}

	return message, nil
}

func GetUserCreated10FromMessage(message Message) (UserCrated10, error) {
	var userCreated UserCrated10

	err := json.Unmarshal(message.Body, &userCreated)
	if err != nil {
		return userCreated, err
	}

	return userCreated, nil
}

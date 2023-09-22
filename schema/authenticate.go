package schema

import "encoding/json"

const (
	CommandIdentifyUser = "identify-user"
)

const (
	CommandIdentifyUser10 = CommandIdentifyUser + "-" + Version10
)

const (
	EventNameUserCreated = "user-created"
)

const (
	PathEventAuthenticateUserCreated10 = PathMsAuthenticate + ".event." + EventNameUserCreated + "-" + Version10
)

type IdentifyUser10Request struct {
	Token string `json:"token"`
}

type IdentifyUser10Response struct {
	ID int64 `json:"id"`
}

type UserCrated10 struct {
	ID   int64  `json:"id"`
	Role string `json:"role"`
}

func GenerateMessageForIdentifyUser10Response(id int64) (Message, error) {
	body := IdentifyUser10Response{
		ID: id,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return Message{}, err
	}

	message := Message{
		Method:  CommandIdentifyUser,
		Version: Version10,
		Body:    bodyBytes,
	}

	return message, nil
}

func GetIdentifyUser10RequestFromMessage(message Message) (IdentifyUser10Request, error) {
	var identifyUser10Request IdentifyUser10Request

	err := json.Unmarshal(message.Body, &identifyUser10Request)
	if err != nil {
		return identifyUser10Request, err
	}

	return identifyUser10Request, nil
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

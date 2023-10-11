package schema

import "encoding/json"

const (
	PathMsAuthenticate = "authenticate"
)

const (
	CommandIdentifyUser  = "identify-user"
	CommandGetDevelopers = "get-developers"
)

const (
	CommandIdentifyUser10  = CommandIdentifyUser + "-" + Version10
	CommandGetDevelopers10 = CommandGetDevelopers + "-" + Version10
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

type GetDevelopers10Request struct {
	FromUserId int64 `json:"from_user_id"`
}

type GetDevelopers10Response struct {
	Developers []Developer `json:"developers"`
}

type Developer struct {
	UserId int64 `json:"user_id"`
}

type UserCrated10 struct {
	ID   int64  `json:"id"`
	Role string `json:"role"`
}

func GenerateDeveloperStruct(userId int64) Developer {
	return Developer{
		UserId: userId,
	}
}

func FillBodyGetDevelopers10ResponseForMessage(developers []Developer, message *Message) error {
	body := GetDevelopers10Response{
		Developers: developers,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	message.Body = bodyBytes

	return nil
}

func FillBodyIdentifyUser10ResponseForMessage(userId int64, message *Message) error {
	body := IdentifyUser10Response{
		ID: userId,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	message.Body = bodyBytes

	return nil
}

func GetIdentifyUser10RequestFromMessage(message Message) (IdentifyUser10Request, error) {
	var identifyUser10Request IdentifyUser10Request

	err := json.Unmarshal(message.Body, &identifyUser10Request)
	if err != nil {
		return identifyUser10Request, err
	}

	return identifyUser10Request, nil
}

func GetGetDevelopers10RequestFromMessage(message Message) (GetDevelopers10Request, error) {
	var getDevelopers10Request GetDevelopers10Request

	err := json.Unmarshal(message.Body, &getDevelopers10Request)
	if err != nil {
		return getDevelopers10Request, err
	}

	return getDevelopers10Request, nil
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

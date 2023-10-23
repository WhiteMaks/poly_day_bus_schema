package schema

import "encoding/json"

const (
	PathMsDeveloperService = "developer-service"
)

const (
	CommandCreateRelease = "create-release"
	CommandDeleteRelease = "delete-release"
	CommandUpdateRelease = "update-release"
	CommandGetReleases   = "get-releases"
)

const (
	CommandCreateRelease10 = CommandCreateRelease + "-" + Version10
	CommandDeleteRelease10 = CommandDeleteRelease + "-" + Version10
	CommandUpdateRelease10 = CommandUpdateRelease + "-" + Version10
	CommandGetReleases10   = CommandGetReleases + "-" + Version10
)

const (
	PathSubscriptionDeveloperServiceAuthenticateUserCreated10 = PathMsDeveloperService + "." + PathEventAuthenticateUserCreated10
)

type CreateRelease10Request struct {
	Name string `json:"name"`
}

type CreateRelease10Response struct {
	ID int64 `json:"id"`
}

type GetReleases10Request struct {
}

type GetReleases10Response struct {
	Releases []Release `json:"releases"`
}

type Release struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type UpdateRelease10Request struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type UpdateRelease10Response struct {
}

type DeleteRelease10Request struct {
	ID int64 `json:"id"`
}

type DeleteRelease10Response struct {
}

func GetUpdateRelease10RequestFromMessage(message Message) (UpdateRelease10Request, error) {
	var updateRelease10Request UpdateRelease10Request

	err := json.Unmarshal(message.Body, &updateRelease10Request)
	if err != nil {
		return updateRelease10Request, err
	}

	return updateRelease10Request, nil
}

func GetDeleteRelease10RequestFromMessage(message Message) (DeleteRelease10Request, error) {
	var deleteRelease10Request DeleteRelease10Request

	err := json.Unmarshal(message.Body, &deleteRelease10Request)
	if err != nil {
		return deleteRelease10Request, err
	}

	return deleteRelease10Request, nil
}

func FillBodyCreateRelease10ResponseForMessage(releaseId int64, message *Message) error {
	body := CreateRelease10Response{
		ID: releaseId,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	message.Body = bodyBytes

	return nil
}

func GetCreateRelease10RequestFromMessage(message Message) (CreateRelease10Request, error) {
	var createRelease10Request CreateRelease10Request

	err := json.Unmarshal(message.Body, &createRelease10Request)
	if err != nil {
		return createRelease10Request, err
	}

	return createRelease10Request, nil
}

func GenerateReleaseStruct(id int64, name string) Release {
	return Release{
		ID:   id,
		Name: name,
	}
}

func FillBodyGetReleases10ResponseForMessage(releases []Release, message *Message) error {
	body := GetReleases10Response{
		Releases: releases,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	message.Body = bodyBytes

	return nil
}
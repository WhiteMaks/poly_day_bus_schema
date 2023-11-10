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

	CommandCreateEntity = "create-entity"
	CommandDeleteEntity = "delete-entity"
	CommandUpdateEntity = "update-entity"
	CommandGetEntities  = "get-entities"
)

const (
	CommandCreateRelease10 = CommandCreateRelease + "-" + Version10
	CommandDeleteRelease10 = CommandDeleteRelease + "-" + Version10
	CommandUpdateRelease10 = CommandUpdateRelease + "-" + Version10
	CommandGetReleases10   = CommandGetReleases + "-" + Version10

	CommandCreateEntity10 = CommandCreateEntity + "-" + Version10
	CommandDeleteEntity10 = CommandDeleteEntity + "-" + Version10
	CommandUpdateEntity10 = CommandUpdateEntity + "-" + Version10
	CommandGetEntities10  = CommandGetEntities + "-" + Version10
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

type EntityType string

const (
	FOLDER EntityType = "FOLDER"
	MODEL  EntityType = "MODEL"
	SHADER EntityType = "SHADER"
	SCRIPT EntityType = "SCRIPT"
	MUSIC  EntityType = "MUSIC"
	SOUND  EntityType = "SOUND"
)

func (et EntityType) String() *string {
	var result string

	switch et {
	case FOLDER:
		{
			result = "FOLDER"
			break
		}
	case MODEL:
		{
			result = "MODEL"
			break
		}
	case SHADER:
		{
			result = "SHADER"
			break
		}
	case SCRIPT:
		{
			result = "SCRIPT"
			break
		}
	case MUSIC:
		{
			result = "MUSIC"
			break
		}
	case SOUND:
		{
			result = "SOUND"
			break
		}
	default:
		return nil
	}

	return &result
}

type ShaderSource struct {
	Vertex   string `json:"vertex"`
	Fragment string `json:"fragment"`
}

type ModelSource struct {
	FileBase64 string `json:"file_base_64"`
}

type Entity[SOURCE any] struct {
	ID       int64            `json:"id"`
	Name     string           `json:"name"`
	Type     EntityType       `json:"type"`
	Source   SOURCE           `json:"source"`
	Entities []Entity[SOURCE] `json:"entities"`
}

type GetEntities10Request struct {
}

type GetEntities10Response struct {
	Entities []Entity[any] `json:"entities"`
}

type CreateEntity10Request struct {
	Name           string     `json:"name"`
	Type           EntityType `json:"type"`
	ParentEntityId int64      `json:"parent_entity_id"`
}

type CreateEntity10Response struct {
	ID int64 `json:"id"`
}

type DeleteEntity10Request struct {
	ID int64 `json:"id"`
}

type DeleteEntity10Response struct {
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

func GetDeleteEntity10RequestFromMessage(message Message) (DeleteEntity10Request, error) {
	var deleteEntity10Request DeleteEntity10Request

	err := json.Unmarshal(message.Body, &deleteEntity10Request)
	if err != nil {
		return deleteEntity10Request, err
	}

	return deleteEntity10Request, nil
}

func GetCreateEntity10RequestFromMessage(message Message) (CreateEntity10Request, error) {
	var createEntity10Request CreateEntity10Request

	err := json.Unmarshal(message.Body, &createEntity10Request)
	if err != nil {
		return createEntity10Request, err
	}

	return createEntity10Request, nil
}

func FillBodyCreateEntity10ResponseForMessage(entityId int64, message *Message) error {
	body := CreateEntity10Response{
		ID: entityId,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	message.Body = bodyBytes

	return nil
}

func GenerateEntityStruct[SOURCE any](id int64, name string, entityType EntityType, source SOURCE) Entity[SOURCE] {
	return Entity[SOURCE]{
		ID:     id,
		Name:   name,
		Type:   entityType,
		Source: source,
	}
}

func FillBodyGetEntities10ResponseForMessage(entities []Entity[any], message *Message) error {
	body := GetEntities10Response{
		Entities: entities,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	message.Body = bodyBytes

	return nil
}

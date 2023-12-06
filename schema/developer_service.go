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

	CommandGetModelSource  = "get-model-source"
	CommandSaveModelSource = "save-model-source"

	CommandGetShaderSource  = "get-shader-source"
	CommandSaveShaderSource = "save-shader-source"

	CommandGetShaderPrograms   = "get-shader-programs"
	CommandCreateShaderProgram = "create-shader-program"
	CommandDeleteShaderProgram = "delete-shader-program"
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

	CommandGetModelSource10  = CommandGetModelSource + "-" + Version10
	CommandSaveModelSource10 = CommandSaveModelSource + "-" + Version10

	CommandGetShaderSource10  = CommandGetShaderSource + "-" + Version10
	CommandSaveShaderSource10 = CommandSaveShaderSource + "-" + Version10

	CommandCreateShaderProgram10 = CommandCreateShaderProgram + "-" + Version10
	CommandDeleteShaderProgram10 = CommandDeleteShaderProgram + "-" + Version10
	CommandGetShaderPrograms10   = CommandGetShaderPrograms + "-" + Version10
)

const (
	PathSubscriptionDeveloperServiceAuthenticateUserCreated10 = PathMsDeveloperService + "." + PathEventAuthenticateUserCreated10
)

type GetShaderPrograms10Request struct {
	ReleaseId int64 `json:"release_id"`
}

type GetShaderPrograms10Response struct {
	ShaderPrograms []ShaderProgram `json:"shader_programs"`
}

type ShaderProgram struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	ShaderEntityId int64  `json:"shader_entity_id"`
}

type DeleteShaderProgram10Request struct {
	ID int64 `json:"id"`
}

type DeleteShaderProgram10Response struct {
}

type CreateShaderProgram10Request struct {
	Name           string `json:"name"`
	ShaderEntityId int64  `json:"shader_entity_id"`
	ReleaseId      int64  `json:"release_id"`
}

type CreateShaderProgram10Response struct {
	ID int64 `json:"id"`
}

type SaveShaderSource10Request struct {
	VertexShader   string `json:"vertex_shader"`
	FragmentShader string `json:"fragment_shader"`
	EntityId       int64  `json:"entity_id"`
}

type SaveShaderSource10Response struct {
	ID int64 `json:"id"`
}

type GetShaderSource10Request struct {
	EntityId int64 `json:"entity_id"`
}

type GetShaderSource10Response struct {
	VertexShader   string `json:"vertex_shader"`
	FragmentShader string `json:"fragment_shader"`
	ID             int64  `json:"id"`
}

type SaveModelSource10Request struct {
	SourceBase64 string `json:"source_base_64"`
	EntityId     int64  `json:"entity_id"`
}

type SaveModelSource10Response struct {
	ID int64 `json:"id"`
}

type GetModelSource10Request struct {
	EntityId int64 `json:"entity_id"`
}

type GetModelSource10Response struct {
	SourceBase64 string `json:"source_base_64"`
	ID           int64  `json:"id"`
}

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

type Entity struct {
	ID       int64      `json:"id"`
	Name     string     `json:"name"`
	Type     EntityType `json:"type"`
	Entities []Entity   `json:"entities"`
}

type GetEntities10Request struct {
}

type GetEntities10Response struct {
	Entities []Entity `json:"entities"`
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

func GenerateShaderProgramStruct(id int64, name string, shaderEntityId int64) ShaderProgram {
	return ShaderProgram{
		ID:             id,
		Name:           name,
		ShaderEntityId: shaderEntityId,
	}
}

func GetGetShaderPrograms10RequestFromMessage(message Message) (GetShaderPrograms10Request, error) {
	var getShaderPrograms10Request GetShaderPrograms10Request

	err := json.Unmarshal(message.Body, &getShaderPrograms10Request)
	if err != nil {
		return getShaderPrograms10Request, err
	}

	return getShaderPrograms10Request, nil
}

func FillBodyGetShaderPrograms10ResponseForMessage(shaderPrograms []ShaderProgram, message *Message) error {
	body := GetShaderPrograms10Response{
		ShaderPrograms: shaderPrograms,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	message.Body = bodyBytes

	return nil
}

func GetDeleteShaderProgram10RequestFromMessage(message Message) (DeleteShaderProgram10Request, error) {
	var deleteShaderProgram10Request DeleteShaderProgram10Request

	err := json.Unmarshal(message.Body, &deleteShaderProgram10Request)
	if err != nil {
		return deleteShaderProgram10Request, err
	}

	return deleteShaderProgram10Request, nil
}

func GetCreateShaderProgram10RequestFromMessage(message Message) (CreateShaderProgram10Request, error) {
	var createShaderProgram10Request CreateShaderProgram10Request

	err := json.Unmarshal(message.Body, &createShaderProgram10Request)
	if err != nil {
		return createShaderProgram10Request, err
	}

	return createShaderProgram10Request, nil
}

func FillBodyCreateShaderProgram10ResponseForMessage(shaderProgramId int64, message *Message) error {
	body := CreateShaderProgram10Response{
		ID: shaderProgramId,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	message.Body = bodyBytes

	return nil
}

func GetGetShaderSource10RequestFromMessage(message Message) (GetShaderSource10Request, error) {
	var getShaderSource10Request GetShaderSource10Request

	err := json.Unmarshal(message.Body, &getShaderSource10Request)
	if err != nil {
		return getShaderSource10Request, err
	}

	return getShaderSource10Request, nil
}

func FillBodyGetShaderSource10ResponseForMessage(sourceId int64, vertexShader string, fragmentShader string, message *Message) error {
	body := GetShaderSource10Response{
		ID:             sourceId,
		VertexShader:   vertexShader,
		FragmentShader: fragmentShader,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	message.Body = bodyBytes

	return nil
}

func GetSaveShaderSource10RequestFromMessage(message Message) (SaveShaderSource10Request, error) {
	var saveShaderSource10Request SaveShaderSource10Request

	err := json.Unmarshal(message.Body, &saveShaderSource10Request)
	if err != nil {
		return saveShaderSource10Request, err
	}

	return saveShaderSource10Request, nil
}

func FillBodySaveShaderSource10ResponseForMessage(sourceId int64, message *Message) error {
	body := SaveShaderSource10Response{
		ID: sourceId,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	message.Body = bodyBytes

	return nil
}

func GetGetModelSource10RequestFromMessage(message Message) (GetModelSource10Request, error) {
	var getModelSource10Request GetModelSource10Request

	err := json.Unmarshal(message.Body, &getModelSource10Request)
	if err != nil {
		return getModelSource10Request, err
	}

	return getModelSource10Request, nil
}

func FillBodyGetModelSource10ResponseForMessage(sourceId int64, sourceBase64 string, message *Message) error {
	body := GetModelSource10Response{
		ID:           sourceId,
		SourceBase64: sourceBase64,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	message.Body = bodyBytes

	return nil
}

func GetSaveModelSource10RequestFromMessage(message Message) (SaveModelSource10Request, error) {
	var saveModelSource10Request SaveModelSource10Request

	err := json.Unmarshal(message.Body, &saveModelSource10Request)
	if err != nil {
		return saveModelSource10Request, err
	}

	return saveModelSource10Request, nil
}

func FillBodySaveModelSource10ResponseForMessage(sourceId int64, message *Message) error {
	body := SaveModelSource10Response{
		ID: sourceId,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	message.Body = bodyBytes

	return nil
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

func GenerateEntityStruct(id int64, name string, entityType EntityType) Entity {
	return Entity{
		ID:   id,
		Name: name,
		Type: entityType,
	}
}

func FillBodyGetEntities10ResponseForMessage(entities []Entity, message *Message) error {
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

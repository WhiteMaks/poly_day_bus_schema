package schema

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetIdentifyUser10RequestFromMessage(t *testing.T) {
	message := Message{
		Method:  CommandIdentifyUser,
		Version: Version10,
		Body:    []byte("{\"token\":\"1234token1234\"}"),
	}

	expectedIdentifyUser10Request := IdentifyUser10Request{
		Token: "1234token1234",
	}

	actualIdentifyUser10Request, err := GetIdentifyUser10RequestFromMessage(message)
	assert.NoError(t, err, "No error must be")

	assert.Equal(t, expectedIdentifyUser10Request, actualIdentifyUser10Request, "Identify user requests must be equals")
}

func TestGenerateMessageForUserCreated10(t *testing.T) {
	expectedMessage := Message{
		Method:  EventNameUserCreated,
		Version: Version10,
		Body:    []byte("{\"id\":1,\"role\":\"ROLE_TEST\"}"),
	}

	actualMessage, err := GenerateMessageForUserCreated10(1, "ROLE_TEST")
	assert.NoError(t, err, "No error must be")

	assert.Equal(t, expectedMessage, actualMessage, "Messages must be equals")
}

func TestGetUserCreated10FromMessage(t *testing.T) {
	message := Message{
		Method:  EventNameUserCreated,
		Version: Version10,
		Body:    []byte("{\"id\":1,\"role\":\"ROLE_TEST\"}"),
	}

	expectedUserCreated10 := UserCrated10{
		ID:   1,
		Role: "ROLE_TEST",
	}

	actualUserCreated10, err := GetUserCreated10FromMessage(message)
	assert.NoError(t, err, "No error must be")

	assert.Equal(t, expectedUserCreated10, actualUserCreated10, "User created events must be equals")
}

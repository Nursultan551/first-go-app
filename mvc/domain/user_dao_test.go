package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserNoUserFound(t *testing.T) {
//	Initialization:

//	Execution:
	user, err := GetUser(0)

//	Validation
	assert.Nil(t, user, "we were not expecting a user with id 0")
	assert.NotNil(t, err)

	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "user 0 does not exists", err.Message)

}

func TestGetUserNoError(t *testing.T) {
//	Initialization:

//	Execution:
	user, err := GetUser(123)

//	Validation
	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Nurik", user.FirstName)
	assert.EqualValues(t, "Kozhomkul", user.LastName)
	assert.EqualValues(t, "user@neon.kg", user.Email)
}

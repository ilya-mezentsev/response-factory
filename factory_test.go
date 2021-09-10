package response_factory

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestDefaultResponse(t *testing.T) {
	r := DefaultResponse()

	assert.True(t, r.IsOk())
	assert.False(t, r.HasData())
	assert.Equal(t, http.StatusNoContent, r.HttpStatus())
	assert.Nil(t, r.Data())
}

func TestSuccessResponse(t *testing.T) {
	someData := `data`
	r := SuccessResponse(someData)

	assert.True(t, r.IsOk())
	assert.True(t, r.HasData())
	assert.Equal(t, http.StatusOK, r.HttpStatus())
	assert.Equal(t, someData, r.Data())
}

func TestServerErrorResponse(t *testing.T) {
	someData := `data`
	r := ServerError(someData)

	assert.False(t, r.IsOk())
	assert.True(t, r.HasData())
	assert.Equal(t, http.StatusInternalServerError, r.HttpStatus())
	assert.Equal(t, someData, r.Data())
}

func TestClientErrorResponse(t *testing.T) {
	someData := `data`
	r := ClientError(someData)

	assert.False(t, r.IsOk())
	assert.True(t, r.HasData())
	assert.Equal(t, http.StatusBadRequest, r.HttpStatus())
	assert.Equal(t, someData, r.Data())
}

func TestEmptyClientError(t *testing.T) {
	r := EmptyClientError()

	assert.False(t, r.IsOk())
	assert.False(t, r.HasData())
	assert.Equal(t, http.StatusBadRequest, r.HttpStatus())
	assert.Nil(t, r.Data())
}

func TestForbiddenErrorResponse(t *testing.T) {
	someData := `data`
	r := ForbiddenError(someData)

	assert.False(t, r.IsOk())
	assert.True(t, r.HasData())
	assert.Equal(t, http.StatusForbidden, r.HttpStatus())
	assert.Equal(t, someData, r.Data())
}

func TestUnauthorizedErrorResponse(t *testing.T) {
	someData := `data`
	r := UnauthorizedError(someData)

	assert.False(t, r.IsOk())
	assert.True(t, r.HasData())
	assert.Equal(t, http.StatusUnauthorized, r.HttpStatus())
	assert.Equal(t, someData, r.Data())
}

func TestEmptyServerError(t *testing.T) {
	r := EmptyServerError()

	assert.False(t, r.IsOk())
	assert.False(t, r.HasData())
	assert.Equal(t, http.StatusInternalServerError, r.HttpStatus())
	assert.Nil(t, r.Data())
}

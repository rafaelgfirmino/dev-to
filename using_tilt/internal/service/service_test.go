package service

import (
	"errors"
	"testing"

	"github.com/rafaelgfirmino/dev_to/using_tilt/internal/domain"
	"github.com/rafaelgfirmino/dev_to/using_tilt/internal/port"
	"github.com/stretchr/testify/assert"
)

// Testando o método GetUserById

func TestGetUserById_Success(t *testing.T) {
	// Criação do mock
	mockRepo := port.NewMockUserDbRepository(t)

	// Configuração da expectativa: quando FindUserById for chamado com "1", deve retornar um usuário
	mockRepo.On("FindUserById", "1").Return(&domain.User{ID: "1", Email: "test@example.com"}, nil)

	// Criação do UserService com o mock
	userService := NewUserService(mockRepo)

	// Chamando o método GetUserById
	user, err := userService.GetUserById("1")

	// Verificando as expectativas
	mockRepo.AssertExpectations(t)

	// Validando os resultados
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "1", user.ID)
	assert.Equal(t, "test@example.com", user.Email)
}

// Testando o método GetUserById quando ocorre um erro no repositório
func TestGetUserById_Error(t *testing.T) {
	// Criação do mock
	mockRepo := port.NewMockUserDbRepository(t)

	// Configuração da expectativa: quando FindUserById for chamado com "1", deve retornar um erro
	mockRepo.On("FindUserById", "1").Return(nil, errors.New("user not found"))

	// Criação do UserService com o mock
	userService := NewUserService(mockRepo)

	// Chamando o método GetUserById
	user, err := userService.GetUserById("1")

	// Verificando as expectativas
	mockRepo.AssertExpectations(t)

	// Validando os resultados
	assert.NotNil(t, err)
	assert.Nil(t, user)
	assert.Equal(t, "user not found", err.Error())
}

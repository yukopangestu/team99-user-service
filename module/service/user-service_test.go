package service

import (
	"errors"
	"testing"
	"time"

	"team99_user_service/module/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository is a mock implementation of UserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetUser(request model.GetUserRequest) ([]model.User, error) {
	args := m.Called(request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.User), args.Error(1)
}

func (m *MockUserRepository) GetUserById(id string) (model.User, error) {
	args := m.Called(id)
	return args.Get(0).(model.User), args.Error(1)
}

func (m *MockUserRepository) CreateUser(user model.User) (model.User, error) {
	args := m.Called(user)
	return args.Get(0).(model.User), args.Error(1)
}

// TestGetAllUser_Success tests successful retrieval of all users
func TestGetAllUser_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	request := model.GetUserRequest{
		PageNum:  1,
		PageSize: 10,
	}

	expectedUsers := []model.User{
		{
			Id:        1,
			Name:      "John Doe",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Id:        2,
			Name:      "Jane Smith",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	mockRepo.On("GetUser", request).Return(expectedUsers, nil)

	result, err := service.GetAllUser(request)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, "John Doe", result[0].Name)
	assert.Equal(t, "Jane Smith", result[1].Name)
	mockRepo.AssertExpectations(t)
}

// TestGetAllUser_RepositoryError tests repository error scenario
func TestGetAllUser_RepositoryError(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	request := model.GetUserRequest{
		PageNum:  1,
		PageSize: 10,
	}

	expectedError := errors.New("database connection failed")
	mockRepo.On("GetUser", request).Return(nil, expectedError)

	result, err := service.GetAllUser(request)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

// TestGetAllUser_EmptyResult tests empty result scenario
func TestGetAllUser_EmptyResult(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	request := model.GetUserRequest{
		PageNum:  1,
		PageSize: 10,
	}

	emptyUsers := []model.User{}
	mockRepo.On("GetUser", request).Return(emptyUsers, nil)

	result, err := service.GetAllUser(request)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 0, len(result))
	mockRepo.AssertExpectations(t)
}

// TestGetUserById_Success tests successful retrieval of user by ID
func TestGetUserById_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	userId := "1"
	expectedUser := model.User{
		Id:        1,
		Name:      "John Doe",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepo.On("GetUserById", userId).Return(expectedUser, nil)

	result, err := service.GetUserById(userId)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser.Id, result.Id)
	assert.Equal(t, expectedUser.Name, result.Name)
	mockRepo.AssertExpectations(t)
}

// TestGetUserById_UserNotFound tests user not found scenario
func TestGetUserById_UserNotFound(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	userId := "999"
	expectedError := errors.New("record not found")

	mockRepo.On("GetUserById", userId).Return(model.User{}, expectedError)

	result, err := service.GetUserById(userId)

	assert.Error(t, err)
	assert.Equal(t, model.User{}, result)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

// TestGetUserById_InvalidId tests invalid ID scenario
func TestGetUserById_InvalidId(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	userId := "invalid"
	expectedError := errors.New("invalid user ID format")

	mockRepo.On("GetUserById", userId).Return(model.User{}, expectedError)

	result, err := service.GetUserById(userId)

	assert.Error(t, err)
	assert.Equal(t, model.User{}, result)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

// TestPostUser_Success tests successful user creation
func TestPostUser_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	request := model.PostUserRequest{
		Name: "New User",
	}

	expectedUser := model.User{
		Id:        1,
		Name:      "New User",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepo.On("CreateUser", mock.MatchedBy(func(user model.User) bool {
		return user.Name == "New User"
	})).Return(expectedUser, nil)

	result, err := service.PostUser(request)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser.Id, result.Id)
	assert.Equal(t, expectedUser.Name, result.Name)
	mockRepo.AssertExpectations(t)
}

// TestPostUser_RepositoryError tests repository error during creation
func TestPostUser_RepositoryError(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	request := model.PostUserRequest{
		Name: "New User",
	}

	expectedError := errors.New("failed to insert user into database")

	mockRepo.On("CreateUser", mock.MatchedBy(func(user model.User) bool {
		return user.Name == "New User"
	})).Return(model.User{}, expectedError)

	result, err := service.PostUser(request)

	assert.Error(t, err)
	assert.Equal(t, model.User{}, result)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

// TestPostUser_DuplicateName tests duplicate user name scenario
func TestPostUser_DuplicateName(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	request := model.PostUserRequest{
		Name: "Existing User",
	}

	expectedError := errors.New("duplicate key value violates unique constraint")

	mockRepo.On("CreateUser", mock.MatchedBy(func(user model.User) bool {
		return user.Name == "Existing User"
	})).Return(model.User{}, expectedError)

	result, err := service.PostUser(request)

	assert.Error(t, err)
	assert.Equal(t, model.User{}, result)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

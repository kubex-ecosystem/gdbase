package tests

// import (
// 	"errors"
// 	"testing"

// 	um "github.com/rafa-mori/gdbase/internal/models/users"
// 	"github.com/rafa-mori/gdbase/types"
// 	"github.com/rafa-mori/xtui/types"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// // MockIUserRepo is a mock implementation of IUserRepo for testing purposes.
// type MockIUserRepo struct {
// 	mock.Mock
// }

// // GetContextDBService implements user.IUserRepo.
// func (m *MockIUserRepo) GetContextDBService() types.DBService {
// 	panic("unimplemented")
// }

// // TableName implements user.IUserRepo.
// func (m *MockIUserRepo) TableName() string {
// 	panic("unimplemented")
// }

// // List implements IUserRepo.
// func (m *MockIUserRepo) List(where ...interface{}) (types.TableHandler, error) {
// 	panic("unimplemented")
// }

// func (m *MockIUserRepo) Close() error {
// 	return nil
// }

// func (m *MockIUserRepo) Create(user um.IUser) (um.IUser, error) {
// 	args := m.Called(user)
// 	return args.Get(0).(um.IUser), args.Error(1)
// }

// func (m *MockIUserRepo) FindOne(args ...interface{}) (um.IUser, error) {
// 	return nil, nil
// }

// func (m *MockIUserRepo) Update(user um.IUser) (um.IUser, error) {
// 	return nil, nil
// }

// func (m *MockIUserRepo) Delete(id string) error {
// 	return nil
// }

// func (m *MockIUserRepo) FindAll(args ...interface{}) ([]um.IUser, error) {
// 	return nil, nil
// }

// // MockIUser is a mock implementation of IUser for testing purposes.
// type MockIUser struct {
// 	mock.Mock
// }

// // TableName implements user.IUser.
// func (m *MockIUser) TableName() string {
// 	panic("unimplemented")
// }

// // SetID implements user.IUser.
// func (m *MockIUser) SetID(id string) {
// 	panic("unimplemented")
// }

// // GetActive implements IUser.
// func (m *MockIUser) GetActive() bool {
// 	panic("unimplemented")
// }

// // GetID implements IUser.
// func (m *MockIUser) GetID() string {
// 	panic("unimplemented")
// }

// // GetName implements IUser.
// func (m *MockIUser) GetName() string {
// 	panic("unimplemented")
// }

// // GetPhone implements IUser.
// func (m *MockIUser) GetPhone() string {
// 	panic("unimplemented")
// }

// // GetRoleID implements IUser.
// func (m *MockIUser) GetRoleID() uint {
// 	panic("unimplemented")
// }

// // SetActive implements IUser.
// func (m *MockIUser) SetActive(active bool) {
// 	panic("unimplemented")
// }

// // SetEmail implements IUser.
// func (m *MockIUser) SetEmail(email string) {
// 	panic("unimplemented")
// }

// // SetName implements IUser.
// func (m *MockIUser) SetName(name string) {
// 	panic("unimplemented")
// }

// // SetPassword implements IUser.
// func (m *MockIUser) SetPassword(password string) error {
// 	panic("unimplemented")
// }

// // SetPhone implements IUser.
// func (m *MockIUser) SetPhone(phone string) {
// 	panic("unimplemented")
// }

// // SetRoleID implements IUser.
// func (m *MockIUser) SetRoleID(roleID uint) {
// 	panic("unimplemented")
// }

// // SetUsername implements IUser.
// func (m *MockIUser) SetUsername(username string) {
// 	panic("unimplemented")
// }

// func (m *MockIUser) GetUsername() string {
// 	args := m.Called()
// 	return args.String(0)
// }

// func (m *MockIUser) GetEmail() string {
// 	args := m.Called()
// 	return args.String(0)
// }

// func (m *MockIUser) GetPassword() string {
// 	args := m.Called()
// 	return args.String(0)
// }

// func TestCreateUser(t *testing.T) {
// 	mockRepo := new(MockIUserRepo)
// 	mockUser := new(MockIUser)

// 	userService := um.NewUserService(mockRepo)

// 	t.Run("should return error when required fields are missing", func(t *testing.T) {
// 		// Set up the mock user model
// 		mockUser.On("GetUsername").Return("")
// 		mockUser.On("GetEmail").Return("")
// 		mockUser.On("GetPassword").Return("")
// 		mockUser.On("GetName").Return("")
// 		mockUser.On("GetPhone").Return("")
// 		mockUser.On("GetRoleID").Return(uint(0))
// 		mockUser.On("GetActive").Return(false)
// 		assert.NotNil(t, mockUser)

// 		// Simulate the error
// 		mockRepo.On("Create", mockUser).Return(nil, errors.New("missing required fields"))

// 		// Call the CreateUser method
// 		// with the mock user model
// 		// and expect it to return an error
// 		_, err := userService.CreateUser(mockUser)

// 		// Assert that the error is not nil
// 		// and the created user is nil
// 		// and the error message is as expected
// 		assert.NotNil(t, err)
// 		// Assert that the error message is as expected
// 		assert.EqualError(t, err, "missing required fields")

// 	})

// 	t.Run("should return error when repo.Create fails", func(t *testing.T) {
// 		mockUser.On("GetUsername").Return("testuser")
// 		mockUser.On("GetEmail").Return("test@example.com")
// 		mockUser.On("GetPassword").Return("password123")
// 		mockRepo.On("Create", mockUser).Return(nil, errors.New("database error"))

// 		createdUser, err := userService.CreateUser(mockUser)

// 		assert.Nil(t, createdUser)
// 		assert.EqualError(t, err, "error creating user: database error")
// 		mockUser.AssertExpectations(t)
// 		mockRepo.AssertExpectations(t)
// 	})

// 	t.Run("should create user successfully", func(t *testing.T) {
// 		mockUser.On("GetUsername").Return("testuser")
// 		mockUser.On("GetEmail").Return("test@example.com")
// 		mockUser.On("GetPassword").Return("password123")
// 		mockRepo.On("Create", mockUser).Return(mockUser, nil)

// 		createdUser, err := userService.CreateUser(mockUser)

// 		assert.NotNil(t, createdUser)
// 		assert.NoError(t, err)
// 		mockUser.AssertExpectations(t)
// 		mockRepo.AssertExpectations(t)
// 	})
// }

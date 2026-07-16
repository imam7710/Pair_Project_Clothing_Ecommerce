package usecase

import (
	"Pair_Project_Clothing_Ecommerce/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository adalah tiruan dari repository asli
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetByEmail(email string) (*domain.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func TestAuthUseCase_Login(t *testing.T) {
	mockRepo := new(MockUserRepository)
	authUC := NewAuthUseCase(mockRepo)

	t.Run("Login Sukses", func(t *testing.T) {
		mockUser := &domain.User{Email: "imam@test.com", Password: "password123", Role: "customer"}
		mockRepo.On("GetByEmail", "imam@test.com").Return(mockUser, nil)

		user, err := authUC.Login("imam@test.com", "password123")

		assert.NoError(t, err)
		assert.Equal(t, "imam@test.com", user.Email)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Login Gagal - Password Salah", func(t *testing.T) {
		mockUser := &domain.User{Email: "imam@test.com", Password: "password123"}
		mockRepo.On("GetByEmail", "imam@test.com").Return(mockUser, nil)

		_, err := authUC.Login("imam@test.com", "salahpass")

		assert.Error(t, err)
		assert.Equal(t, "invalid credentials", err.Error())
	})
}

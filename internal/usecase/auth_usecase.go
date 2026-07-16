package usecase

import (
	"Pair_Project_Clothing_Ecommerce/internal/domain"
	"errors"
	"strings"
)

type authUC struct{ repo domain.UserRepository }

func NewAuthUseCase(r domain.UserRepository) domain.AuthUseCase { return &authUC{repo: r} }

func (u *authUC) Login(e, p string) (*domain.User, error) {
	if e == "" || p == "" {
		return nil, errors.New("email dan password wajib diisi")
	}
	return u.repo.Login(e, p)
}

type userUC struct{ repo domain.UserRepository }

func NewUserUseCase(r domain.UserRepository) domain.UserUseCase { return &userUC{repo: r} }

func (u *userUC) GetAccountList() ([]domain.UserReportDTO, error) {
	return u.repo.GetAll()
}

func (u *userUC) RegisterNewAccount(req domain.CreateUserRequest) error {
	req.Email = strings.TrimSpace(req.Email)
	if req.Email == "" || req.Password == "" || req.FullName == "" {
		return errors.New("email, password, dan nama lengkap wajib diisi")
	}
	if req.Role != "admin" && req.Role != "customer" {
		req.Role = "customer"
	}
	return u.repo.Create(req)
}

func (u *userUC) ChangeUserRole(userID int, newRole string) error {
	if userID <= 0 {
		return errors.New("ID user tidak valid")
	}
	if newRole != "admin" && newRole != "customer" {
		return errors.New("role hanya boleh 'admin' atau 'customer'")
	}
	return u.repo.UpdateRole(userID, newRole)
}

func (u *userUC) RemoveAccount(userID int) error {
	if userID <= 0 {
		return errors.New("ID user tidak valid")
	}
	return u.repo.Delete(userID)
}

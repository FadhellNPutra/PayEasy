package usecase

import (
	"errors"
	"fmt"
	"payeasy/entity"
	"payeasy/repository"
	"payeasy/shared/model"
)

type UsersUseCase interface {
	FindUsersByID(id string) (entity.Users, error)
	FindUsersByEmail(email string) (entity.Users, error)
	FindUsersForLogin(username, password string) (entity.Users, error)
	// FindUsersForLogout(user *entity.Users) (entity.Users, error)
	RegisterNewUsers(payload entity.Users) (entity.Users, error)
	UpdateUsers(payload entity.Users) (entity.Users, error)
	ListAll(page, size int) ([]entity.Users, model.Paging, error)
	DeleteUsers(id string) error
}

type usersUseCase struct {
	repo repository.UsersRepository
}

// FindUsersForLogout implements UsersUseCase.
// func (u *usersUseCase) FindUsersForLogout(email, password entity.Users) (entity.Users, error) {
// 	return u.repo.GetUsersForLogin(email, password)
// }

func (u *usersUseCase) DeleteUsers(id string) error {
	if _, err := u.repo.GetUsersById(id); err != nil {
		return err
	}

	return u.repo.DeleteUser(id)
}

// FindUsersByID implements UsersUseCase.
func (u *usersUseCase) FindUsersByID(id string) (entity.Users, error) {
	if id == "" {
		return entity.Users{}, errors.New("id harus diisi")
	}
	return u.repo.GetUsersById(id)
}

// FindUsersByUsername implements UsersUseCase.
func (u *usersUseCase) FindUsersByEmail(email string) (entity.Users, error) {
	if email == "" {
		return entity.Users{}, errors.New("email harus diisi")
	}
	return u.repo.GetUsersByEmail(email)
}

// FindUsersForLogin implements UsersUseCase.
func (u *usersUseCase) FindUsersForLogin(email string, password string) (entity.Users, error) {
	if email == "" || password == "" {
		return entity.Users{}, errors.New("email harus diisi")
	}
	return u.repo.GetUsersForLogin(email, password)
}

// ListAll implements UsersUseCase.
func (u *usersUseCase) ListAll(page int, size int) ([]entity.Users, model.Paging, error) {
	return u.repo.List(page, size)
}

// RegisterNewUsers implements UsersUseCase.
func (u *usersUseCase) RegisterNewUsers(payload entity.Users) (entity.Users, error) {
	if payload.Name == "" || payload.Email == "" || payload.Password == "" || payload.Address == "" || payload.Role == "" || payload.Balance == 0 {
		return entity.Users{}, fmt.Errorf("oops, field required")
	}

	users, err := u.repo.CreateUsers(payload)
	if err != nil {
		return entity.Users{}, fmt.Errorf("oppps, failed to save data users :%v", err.Error())
	}
	return users, nil
}

// UpdateUsers implements UsersUseCase.
func (u *usersUseCase) UpdateUsers(payload entity.Users) (entity.Users, error) {
	if payload.ID == "" || payload.Name == "" || payload.Email == "" || payload.Password == "" || payload.Address == "" || payload.Role == "" || payload.Balance == 0 {
		return entity.Users{}, fmt.Errorf("oops, field required")
	}

	users, err := u.repo.UpdateUsers(payload)
	if err != nil {
		return entity.Users{}, fmt.Errorf("oppps, failed to save data users :%v", err.Error())
	}
	return users, nil
}

func NewUsersUseCase(repo repository.UsersRepository) UsersUseCase {
	return &usersUseCase{repo: repo}
}

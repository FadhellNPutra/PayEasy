package repository

import (
	"database/sql"
	"log"
	"math"
	"payeasy/config"
	"payeasy/entity"
	"payeasy/shared/model"
	"time"
)

type UsersRepository interface {
	CreateUsers(payload entity.Users) (entity.Users, error)
	GetUsersById(id string) (entity.Users, error)
	GetUsersByEmail(email string) (entity.Users, error)
	GetUsersForLogin(email, password string) (entity.Users, error)
	UpdateUsers(payload entity.Users) (entity.Users, error)
	List(page, size int) ([]entity.Users, model.Paging, error)
	DeleteUser(id string) error
}

type usersRepository struct {
	db *sql.DB
}

// DeleteUser implements UsersRepository.
func (u *usersRepository) DeleteUser(id string) error {
	err := u.db.QueryRow(config.DeleteUser, id).Err()
	if err != nil {
		log.Println("usersRepository.DeleteUser.QueryRow: ", err.Error())
		return err
	}

	log.Println(err)
	return nil
}

// CreateUsers implements UsersRepository.
func (u *usersRepository) CreateUsers(payload entity.Users) (entity.Users, error) {
	var users entity.Users

	payload.UpdatedAt = time.Now()

	err := u.db.QueryRow(config.InsertUser,
		payload.Name,
		payload.Email,
		payload.Password,
		payload.Number,
		payload.Address,
		payload.Role,
		payload.Balance,
	).Scan(&users.ID, &users.CreatedAt, &users.UpdatedAt)

	if err != nil {
		log.Println("usersRepository.QueryRow: ", err.Error())
		return entity.Users{}, err
	}

	users.Name = payload.Name
	users.Email = payload.Email
	users.Password = payload.Password
	users.Number = payload.Number
	users.Address = payload.Address
	users.Role = payload.Role
	users.Balance = payload.Balance

	return users, nil
}

// GetUsersByEmail implements UsersRepository.
func (u *usersRepository) GetUsersByEmail(email string) (entity.Users, error) {
	var users entity.Users
	err := u.db.QueryRow(config.SelectUserByEmail, email).Scan(
		&users.ID,
		&users.Name,
		&users.Email,
		&users.Password,
		&users.Number,
		&users.Address,
		&users.Role,
		&users.Balance,
		&users.CreatedAt,
		&users.UpdatedAt)
	if err != nil {
		log.Println("usersRepository.GetUsersByEmail.QueryRow: ", err.Error())
		return entity.Users{}, err
	}
	return users, nil
}

// GetUsersById implements UsersRepository.
func (u *usersRepository) GetUsersById(id string) (entity.Users, error) {
	var users entity.Users
	err := u.db.QueryRow(config.SelectUserById, id).Scan(
		&users.ID,
		&users.Name,
		&users.Email,
		&users.Password,
		&users.Number,
		&users.Address,
		&users.Role,
		&users.Balance,
		&users.CreatedAt,
		&users.UpdatedAt)
	if err != nil {
		log.Println("usersRepository.GetUsersByID.QueryRow: ", err.Error())
		return entity.Users{}, err
	}
	return users, nil
}
func (u *usersRepository) GetUsersOnlyId(id string) (entity.Users, error) {
	var users entity.Users
	err := u.db.QueryRow(config.SelectUserOnlyId, id).Scan(
		&users.ID)
	if err != nil {
		log.Println("usersRepository.GetUsersByID.QueryRow: ", err.Error())
		return entity.Users{}, err
	}
	return users, nil
}

// GetUsersForLogin implements UsersRepository.
func (u *usersRepository) GetUsersForLogin(email string, password string) (entity.Users, error) {
	var users entity.Users
	err := u.db.QueryRow(config.SelectUserForLogin, email, password).Scan(
		&users.ID,
		&users.Name,
		&users.Email,
		&users.Password,
		&users.Role)
	if err != nil {
		log.Println("usersRepository.GetUsersForLogin.QueryRow: ", err.Error())
		return entity.Users{}, err
	}
	return users, nil
}

// List implements UsersRepository.
func (u *usersRepository) List(page int, size int) ([]entity.Users, model.Paging, error) {
	var users []entity.Users
	offset := (page - 1) * size
	rows, err := u.db.Query(config.SelectAllUser, size, offset)
	if err != nil {
		log.Println("usersRepository.Query:", err.Error())
		return nil, model.Paging{}, err
	}
	for rows.Next() {
		var user entity.Users
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.Number,
			&user.Address,
			&user.Role,
			&user.Balance,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			log.Println("usersRepository.Rows.Next():", err.Error())
			return nil, model.Paging{}, err
		}

		users = append(users, user)
	}

	totalRows := 0
	if err := u.db.QueryRow("SELECT COUNT(*) FROM users").Scan(&totalRows); err != nil {
		return nil, model.Paging{}, err
	}

	paging := model.Paging{
		Page:        page,
		RowsPerPage: size,
		TotalRows:   totalRows,
		TotalPages:  int(math.Ceil(float64(totalRows) / float64(size))),
	}
	return users, paging, nil
}

// UpdateUsers implements UsersRepository.
func (u *usersRepository) UpdateUsers(payload entity.Users) (entity.Users, error) {
	var users entity.Users

	payload.UpdatedAt = time.Now()

	err := u.db.QueryRow(config.UpdateUser,
		payload.Name,
		payload.Email,
		payload.Password,
		payload.Number,
		payload.Address,
		payload.Role,
		payload.Balance,
		payload.ID,
	).Scan(&users.CreatedAt, &users.UpdatedAt)

	if err != nil {
		log.Println("usersRepository.QueryRow: ", err.Error())
		return entity.Users{}, err
	}

	users.ID = payload.ID
	users.Name = payload.Name
	users.Email = payload.Email
	users.Password = payload.Password
	users.Number = payload.Number
	users.Address = payload.Address
	users.Role = payload.Role
	users.Balance = payload.Balance

	return users, nil
}

func NewUsersRepository(db *sql.DB) UsersRepository {
	return &usersRepository{db: db}
}

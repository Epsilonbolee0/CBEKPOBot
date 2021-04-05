package users

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type UserRepositoryInterface interface {
	CreateNewUser(telegramID string) error

	GetAllUsers() ([]User, error)
	GetUserByID(telegramID string) (User, error)

	UpdateUserFromTime(telegramID string, newFromTime time.Time) error
	UpdateUserToTime(telegramID string, newToTime time.Time) error
	UpdateUserName(telegramID string, newName string) error

	DeleteAllUsers() error
	DeleteUserByID(telegramID string) error
}

type UserRepository struct {
	Database *gorm.DB
}

var userRepository *UserRepository

func initUserRepository() (*UserRepository, error) {
	dsn := "host=localhost user=postgres password=1 dbname=svekrov_bot port=5432 TimeZone=Europe/Moscow"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&User{})

	return &UserRepository{db}, nil
}

func GetUserRepository() *UserRepository {
	if userRepository == nil {
		instance, err := initUserRepository()
		if err != nil {
			panic(err)
		} else {
			userRepository = instance
		}
	}

	return userRepository
}

func (ur *UserRepository) CreateNewUser(telegramID string) error {
	db := ur.Database
	newUser := User{ID: telegramID}

	status := db.Create(&newUser)

	return status.Error
}

func (ur *UserRepository) GetAllUsers() []User {
	var allUsers []User
	db := ur.Database

	status := db.Find(&allUsers)
	if status.Error != nil {
		panic(status.Error)
	}

	return allUsers
}

func (ur *UserRepository) GetUserByID(telegramID string) *User {
	var user User
	db := ur.Database

	status := db.Find(&user).Where("id = ?", telegramID)
	if status.Error != nil {
		panic(status.Error)
	}

	return &user
}

func (ur *UserRepository) UpdateUserFromTime(telegramID string, newFromTime time.Time) error {
	db := ur.Database
	user := User{ID: telegramID}

	status := db.Model(user).Update("MessageFromTime", newFromTime)

	return status.Error
}

func (ur *UserRepository) UpdateUserToTime(telegramID string, newToTime time.Time) error {
	db := ur.Database
	user := User{ID: telegramID}

	status := db.Model(user).Update("MessageFromTime", newToTime)

	return status.Error
}

func (ur *UserRepository) UpdateUserName(telegramID string, newName string) error {
	db := ur.Database
	user := User{ID: telegramID}

	status := db.Model(user).Update("Name", newName)

	return status.Error
}

func (ur *UserRepository) DeleteAllUsers() error {
	db := ur.Database

	status := db.Exec("DELETE FROM users")

	return status.Error
}

func (ur *UserRepository) DeleteUserByID(telegramID string) error {
	db := ur.Database

	status := db.Where("id = ?", telegramID).Delete(&User{})

	return status.Error
}

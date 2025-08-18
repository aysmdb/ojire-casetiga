package models

import (
	"github.com/aysmdb/ojire-casetiga/pkg/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"uniqueIndex"`
	Password string
	Name     string
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GetUserByEmailAndPassword(login LoginRequest) (User, error) {
	db := database.DBConn
	var user User
	err := db.Where("email = ? AND password = ?", login.Email, login.Password).First(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func GetUserByID(id uint) (User, error) {
	db := database.DBConn
	var user User
	err := db.First(&user, id).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func Seeduser() error {
	users := []User{
		{
			Email:    "user1@email.com",
			Password: "123",
			Name:     "User One",
		},
		{
			Email:    "user2@email.com",
			Password: "123",
			Name:     "User Two",
		},
	}

	err := database.DBConn.Transaction(func(tx *gorm.DB) error {
		if err := tx.Unscoped().Where("name is not null").Delete(&User{}).Error; err != nil {
			return err
		}

		if err := tx.Create(&users).Error; err != nil {
			return err
		}

		return nil
	})

	return err
}

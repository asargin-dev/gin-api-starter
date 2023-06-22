package models

import (
	"errors"
	"github.com/asargin-dev/soundproof-backend-demo/pkg/tokenizer"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email         string `gorm:"size:255;not null;unique" json:"email" binding:"required"`
	Password      string `gorm:"size:255;not null;" json:"password" binding:"required"`
	PublicAddress string `gorm:"size:255;not null;" json:"public_address"`
}

func (user *User) CheckIfRegistered() bool {

	var userCheck int64
	DB.Model(&user).Where("email = ?", user.Email).Count(&userCheck)
	return userCheck > 0

}

func (user *User) Register() (*User, error) {

	var err error
	err = DB.Create(user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (user *User) Login() (string, error) {

	err := DB.Model(&user).Where("email = ? AND password = ?", user.Email, user.Password).Take(&user).Error
	if err != nil {
		return "", err
	}

	token, err := tokenizer.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (user *User) GetUserProfile(userId uint) (*User, error) {

	if err := DB.First(&user, userId).Error; err != nil {
		return user, errors.New("No user found")
	}

	return user, nil
}

func (user *User) UpdateUserProfile(publicAddress string) (*User, error) {

	if err := DB.Model(&user).Update("public_address", publicAddress).Error; err != nil {
		return user, errors.New("Update has been failed")
	}

	return user, nil
}

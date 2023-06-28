package repository

import (
	"kmipn-2023/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(email string) (model.User, error)
	CreateUser(user model.User) (model.User, error)
	UpdateUser(user model.User) (model.User, error)
	DeleteUser(user model.User) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByEmail(email string) (model.User, error) {
	var user model.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return model.User{}, nil
	}
	return user, nil
}

func (r *userRepository) CreateUser(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) UpdateUser(user model.User) (model.User, error) {
	if err := r.db.Save(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *userRepository) DeleteUser(user model.User) (model.User, error) {
	if err := r.db.Delete(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

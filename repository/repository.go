package repository

import (
	"goauth/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type RepositoryInterface interface {
	CreateUser(data model.Users) error
	FindByEmail(email string) (model.Users, error)
}

func NewRepository(db *gorm.DB) RepositoryInterface {
	return &UserRepository{
		db: db,
	}
}

// CreateUser implements RepositoryInterface.
func (ur *UserRepository) CreateUser(data model.Users) error {
	err := ur.db.Create(&data)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

// FindByEmail implements RepositoryInterface.
func (ur *UserRepository) FindByEmail(email string) (model.Users, error) {
	var user model.Users
	result := ur.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return model.Users{}, result.Error
	}
	return user, nil
}

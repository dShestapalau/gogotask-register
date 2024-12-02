package repository

import (
	"log"

	"github.com/dshestapalau/gogotask/register/internal/persistence/entity"
	"gorm.io/gorm"
)

type (
	UserCredentialsRepository interface {
		Save(userCredentials entity.UserCredentials) entity.UserCredentials
		FindByLogin(login string) entity.UserCredentials
	}

	userCredentialsRepository struct {
		db *gorm.DB
	}
)

func NewUserCredentialsRepository(db *gorm.DB) UserCredentialsRepository {
	return &userCredentialsRepository{
		db: db,
	}
}

func (repository *userCredentialsRepository) Save(userCredentials entity.UserCredentials) entity.UserCredentials {
	result := repository.db.Create(&userCredentials)
	if result.Error != nil {
		log.Fatalf("Failed to save user_credentials: %v", result.Error)
		return entity.UserCredentials{}
	}

	return userCredentials
}

func (repository *userCredentialsRepository) FindByLogin(login string) entity.UserCredentials {
	var user entity.UserCredentials
	result := repository.db.Debug().Where("login = ?", login).First(&user)

	if result.Error != nil {
		log.Fatalf("Failed to get user_credentials: %v", result.Error)
		return entity.UserCredentials{}
	}

	return user
}

package repository

import (
	"log"

	"github.com/dshestapalau/gogotask/register/internal/persistence/entity"
	"gorm.io/gorm"
)

type (
	UserCredentialsRepository interface {
		Save(userCredentials entity.UserCredentials) entity.UserCredentials
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
	if err := repository.db.Debug().Create(&userCredentials); err != nil {
		log.Fatalf("Failed to save data", err)
		return entity.UserCredentials{}
	}

	return userCredentials
}

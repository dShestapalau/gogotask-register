package service

import (
	"github.com/dshestapalau/gogotask/register/internal/domain/model"
	"github.com/dshestapalau/gogotask/register/internal/persistence/entity"
	"github.com/dshestapalau/gogotask/register/internal/persistence/repository"
)

type (
	UserService interface {
		CreateUser(userCredentials model.UserCredentials) model.UserCredentials
	}

	userService struct {
		userCredentialsRepository repository.UserCredentialsRepository
	}
)

func NewUserService(userCredentialsRepository repository.UserCredentialsRepository) UserService {
	return &userService{
		userCredentialsRepository: userCredentialsRepository,
	}
}

func (service *userService) CreateUser(userCredentials model.UserCredentials) model.UserCredentials {

	savedUserProfile := entity.UserProfile{
		Locale: "ru",
		Status: "CREATED",
	}

	credentials := entity.UserCredentials{
		Login:       userCredentials.Login,
		Password:    userCredentials.Password,
		UserProfile: savedUserProfile,
	}

	service.userCredentialsRepository.Save(credentials)

	return model.UserCredentials{}
}

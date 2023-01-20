package services

import "github.com/najibjodiansyah/gin-users-api/models"

type UserService interface {
	Get() ([]*models.User, error)
	GetByUser(name string) (*models.User, error)
	Create(user *models.User) error
	Update(name string, user *models.User) error
	Delete(name string) error
}

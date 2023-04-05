package repositories

import (
	"dumbflix/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.Users, error)
	GetUser(ID int) (models.Users, error)
	CreateUser(user models.Users) (models.Users, error)
	UpdateUser(user models.Users) (models.Users, error)
	DeleteUser(user models.Users) (models.Users, error)
}

type repository struct {
	db *gorm.DB
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUsers() ([]models.Users, error) {
	var users []models.Users
	err := r.db.Find(&users).Error
	return users, err
}

func (r *repository) GetUser(ID int) (models.Users, error) {
	var user models.Users
	err := r.db.First(&user, ID).Error
	return user, err
}

func (r *repository) CreateUser(user models.Users) (models.Users, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) UpdateUser(user models.Users) (models.Users, error) {
	err := r.db.Save(&user).Error

	return user, err
}

func (r *repository) DeleteUser(user models.Users) (models.Users, error) {
	err := r.db.Delete(&user).Scan(&user).Error

	return user, err
}

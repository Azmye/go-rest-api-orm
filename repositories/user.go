package repositories

import (
	"dumbflix/models"
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.Users, error)
	GetUser(ID int) (models.Users, error)
	CreateUser(user models.Users) (models.Users, error)
	UpdateUser(user models.Users, ID int) (models.Users, error)
	DeleteUser(user models.Users, ID int) (models.Users, error)
}

type repository struct {
	db *gorm.DB
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUsers() ([]models.Users, error) {
	var users []models.Users
	err := r.db.Raw("SELECT * FROM users").Scan(&users).Error
	return users, err
}

func (r *repository) GetUser(ID int) (models.Users, error) {
	var user models.Users
	err := r.db.Raw("SELECT * FROM users WHERE id=?", ID).Scan(&user).Error
	return user, err
}

func (r *repository) CreateUser(user models.Users) (models.Users, error) {
	err := r.db.Exec("INSERT INTO users(name,email,password,created_at, updated_at) VALUES (?,?,?,?,?)", user.Name, user.Email, user.Password, time.Now(), time.Now()).Error

	return user, err
}

func (r *repository) UpdateUser(user models.Users, ID int) (models.Users, error) {
	err := r.db.Raw("UPDATE users SET name=?, email=?, password=?, updated_at=? WHERE id=?", user.Name, user.Email, user.Password, time.Now(), ID).Scan(&user).Error

	return user, err
}

func (r *repository) DeleteUser(user models.Users, ID int) (models.Users, error) {
	err := r.db.Raw("DELETE FROM users WHERE id=?", ID).Scan(&user).Error

	return user, err
}

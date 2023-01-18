package repository

import (
	"campaign-api/domain"
	"gorm.io/gorm"
)

type Repository interface {
	SAVE(user domain.User) (domain.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepositroy(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) SAVE(user domain.User) (domain.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

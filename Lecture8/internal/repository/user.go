package repository

import (
	"errors"
	"fmt"
	"hw8/internal/entity"
	"strings"
	"time"
)

func (r *Repo) GetUserByIdWithPreload(id string) (*entity.User, error) {
	var user entity.User
	if err := r.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, errors.New("the user belonging to this token no logger exists")
	}
	return &user, nil
}
func (r *Repo) GetUserById(id string) (*entity.User, error) {
	var user entity.User
	if err := r.DB.Where("id = ?", fmt.Sprint(id)).First(&user).Error; err != nil {
		return nil, errors.New("the user belonging to this token no logger exists")
	}
	return &user, nil
}

func (r *Repo) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := r.DB.Preload("Role").Where("email = ?", strings.ToLower(email)).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repo) GetUsersByRole(roleId uint) (*[]entity.User, error) {
	var users []entity.User
	if err := r.DB.Preload("Role").Where("role_id = ? ", roleId).Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func (r *Repo) GetUserByResetToken(resetToken string) (*entity.User, error) {
	var user entity.User
	if err := r.DB.Where("password_reset_token = ? AND password_reset_at > ?", resetToken, time.Now()).First(&user).Error; err != nil {
		return nil, errors.New("the reset token is invalid or has expired")
	}
	return &user, nil
}

func (r *Repo) GetUserByVerificationCode(code string) (*entity.User, error) {
	var user entity.User
	if err := r.DB.Where("verification_code = ?", code).First(&user).Error; err != nil {
		return nil, errors.New("invalid verification code or user doesn't exists")
	}
	return &user, nil
}

func (r *Repo) CreateUser(user *entity.User) error {
	err := r.DB.Create(user).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique") {
			return errors.New("user with that email already exists")
		}
		return err
	}
	return nil
}

func (r *Repo) SaveUser(user *entity.User) {
	r.DB.Save(user)
}

package repository

import (
	"context"

	"github.com/Hulhay/aino-skill-test-be/model"
	"gorm.io/gorm"
)

type repository struct {
	qry *gorm.DB
}

type UserRepository interface {
	InsertUser(ctx context.Context, params *model.User) error
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	UpdateStatusLoginTrue(ctx context.Context, email string) error
	UpdateStatusLoginFalse(ctx context.Context, email string) error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &repository{
		qry: db,
	}
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user *model.User

	if err := r.qry.Model(&user).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repository) InsertUser(ctx context.Context, params *model.User) error {
	var user *model.User

	if err := r.qry.Model(&user).Create(map[string]interface{}{
		"name":     params.Name,
		"email":    params.Email,
		"password": params.Password,
		"is_login": params.IsLogin,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) UpdateStatusLoginTrue(ctx context.Context, email string) error {
	var user *model.User

	if err := r.qry.Model(&user).Where("email = ?", email).Updates(map[string]interface{}{
		"is_login": true,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateStatusLoginFalse(ctx context.Context, email string) error {
	var user *model.User

	if err := r.qry.Model(&user).Where("email = ?", email).Updates(map[string]interface{}{
		"is_login": false,
	}).Error; err != nil {
		return err
	}

	return nil
}

package users

import (
	"context"
	"gorm.io/gorm"
	"thegraduate-server/entities"
	"thegraduate-server/interfaces"
)

type UserRepository struct {
	db *gorm.DB
}

func NewStaffRepository(db *gorm.DB) interfaces.IUserRepository {

	return &UserRepository{db: db}
}

func (u *UserRepository) FindAll(ctx context.Context) []entities.UserEntity {

	var result []entities.UserEntity
	err := u.db.WithContext(ctx).Table("users").Find(&result).Error

	if err != nil {
		panic(err)
	}
	return result
}

func (u *UserRepository) Insert(ctx context.Context, entity entities.UserEntity) error {

	err := u.db.WithContext(ctx).Table("users").Create(&entity).Error

	if err != nil {
		return err
	}

	return nil

}

func (u *UserRepository) FindOne(ctx context.Context, username string) (error, entities.UserEntity) {
	var result entities.UserEntity
	err := u.db.Table("users").WithContext(ctx).Where("username=?", username).First(&result).Error

	if err != nil {
		return err, entities.UserEntity{}
	}

	return nil, result

}

func (u *UserRepository) InsertSession(ctx context.Context, data entities.SessionEntity) error {
	//TODO implement me
	err := u.db.Table("Session").WithContext(ctx).Create(&data).Error

	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) FindOneSessionByUsername(ctx context.Context, username string) (entities.SessionEntity, error) {
	//TODO implement me
	var result entities.SessionEntity
	err := u.db.Table("Session").WithContext(ctx).Where("username=?", username).First(&result).Error
	if err != nil {
		return entities.SessionEntity{}, err

	}
	return result, nil

}

func (u *UserRepository) FindByEmail(ctx context.Context, email string) (entities.UserEntity, error) {
	//TODO implement me

	var result entities.UserEntity
	err := u.db.Table("users").WithContext(ctx).Where("email=?", email).First(&result).Error

	if err != nil {
		return entities.UserEntity{}, err
	}
	return result, nil

}

func (u *UserRepository) ChangePassword(ctx context.Context, password string, username string) error {
	err := u.db.Table("users").WithContext(ctx).Where("username=?", username).Update("password", password).Error

	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) DeleteSession(ctx context.Context, username string) error {
	//TODO implement me

	err := u.db.Table("Session").WithContext(ctx).Where("username=?", username).Delete(&entities.SessionEntity{}).Error
	if err != nil {
		return err
	}
	return nil
}

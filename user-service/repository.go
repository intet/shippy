package main

import (
	pb "github.com/intet/shippy/user-service/proto/user-service"
	"github.com/jinzhu/gorm"
)

type Repository interface {
	GetAll(ctx context.Context) ([]*User, error)
	Get(ctx context.Context, id string) (*User, error)
	Create(ctx context.Context, user *pb.User) error
	GetByEmailAndPassword(ctx context.Context, user *User) (*User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func (repo *UserRepository) GetAll(ctx context.Context) ([]*User, error) {
	var users []*User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) Get(ctx context.Context, id string) (*User, error) {
	var user *User
	user.Id = id
	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetByEmailAndPassword(user *pb.User) (*pb.User, error) {
	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) Create(user *pb.User) error {
	if err := repo.db.Create(user).Error; err != nil {
		return err
	}
}
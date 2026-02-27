package database

import (
	"errors"
	"http-demo/models"

	"gorm.io/gorm"
)

type IUserDB interface {
	Create(user *models.User) (*models.User, error)
	GetById(id int) (*models.User, error)
	GetAll() ([]models.User, error)
	DeleteById(id int) (int, error)
}

type UserDB struct {
	DB *gorm.DB
}

func NewUserDB(db *gorm.DB) IUserDB {
	return &UserDB{db}
}

func (ud *UserDB) Create(user *models.User) (*models.User, error) {
	ud.DB.AutoMigrate(&models.User{}) // automatically created as a table

	tx := ud.DB.Create(user)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return user, nil
}

func (ud *UserDB) GetById(id int) (*models.User, error) {
	user := new(models.User)
	tx := ud.DB.First(user, id)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

func (ud *UserDB) GetAll() ([]models.User, error) {
	users := make([]models.User, 0)
	tx := ud.DB.Find(&users)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil
}

func (ud *UserDB) DeleteById(id int) (int, error) {
	tx := ud.DB.Delete(&models.User{}, id)
	if tx.Error != nil {
		return -1, tx.Error
	}

	if tx.RowsAffected > 0 {
		return int(tx.RowsAffected), nil
	} else {
		return -1, errors.New("nothing delete")
	}
}

// func (ud *UserDB) UpdateById(id int, m map[string]any) (*models.User, error) {

// }

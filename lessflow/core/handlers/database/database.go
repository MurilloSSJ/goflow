package database

import (
	"gorm.io/gorm"
)

type Database interface {
	GetConnection() (*gorm.DB, error)
	CloseConnection()
	Insert(interface {}) error
	Update(interface {}) error
	Delete(interface {}) error
	Get(interface {}) error
	GetAll(interface {}) error
}
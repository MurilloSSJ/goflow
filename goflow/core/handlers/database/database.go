package database

type Database interface {
	GetConnection() error
	CloseConnection()
	Insert(interface {}) error
	Update(interface {}) error
	Delete(interface {}) error
	Get(interface {}) error
	GetAll(interface {}) error
}
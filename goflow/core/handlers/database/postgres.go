package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	Connection *gorm.DB
}

func (p *Postgres) GetConnection() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		p.Host, p.Port, p.Username, p.Database, p.Password)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	p.Connection = db
	return nil
}

func (p *Postgres) CloseConnection() {
	sqlDB, err := p.Connection.DB()
	if err != nil {
		return
	}
	sqlDB.Close()
}

func (p *Postgres) Insert(data interface {}) error {
	return p.Connection.Create(data).Error
}

func (p *Postgres) Update(data interface {}) error {
	return p.Connection.Save(data).Error
}

func (p *Postgres) Delete(data interface {}) error {
	return p.Connection.Delete(data).Error
}

func (p *Postgres) Get(data interface {}) error {
	return p.Connection.First(data).Error
}

func (p *Postgres) GetAll(data interface {}) error {
	return p.Connection.Find(data).Error
}


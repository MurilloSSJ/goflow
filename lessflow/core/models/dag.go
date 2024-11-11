package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Dag struct {
	gorm.Model
	DagId string `gorm:"primaryKey"`
	Path string `gorm:"unique"`
	Description string
	Schedule string
	StartDate *time.Time
	EndDate *time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Concurrency int
	MaxActiveTasks int
	TimeOut int
	Catchup bool
	Tags pq.StringArray `gorm:"type:text[]"`
}
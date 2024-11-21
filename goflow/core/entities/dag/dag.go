package DAGModel

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type DAG struct {
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
	Nodes map[string]*Node
}

type DagParams struct {
	DagId string
	Path string
	Description string
	Schedule string
	StartDate *time.Time
	EndDate *time.Time
	Concurrency int
	maxActiveTasks int
	timeOut int
	catchup bool
	tags []string

}

func Constructor(dagParams DagParams) *DAG {
	return &DAG{
		DagId: dagParams.DagId,
		Description: dagParams.Description,
		Schedule: dagParams.Schedule,
		StartDate: dagParams.StartDate,
		EndDate: dagParams.EndDate,
		Concurrency: dagParams.Concurrency,
		MaxActiveTasks: dagParams.maxActiveTasks,
		TimeOut: dagParams.timeOut,
		Catchup: dagParams.catchup,
		Tags: dagParams.tags,
	}
}

func (dag *DAG) AddNode(action func() error) error {
	dag.Nodes[GetNodeName(action)] = &Node{
		Name:         GetNodeName(action),
		Dependencies: nil,
		Action:       action,
	}
	return nil
}


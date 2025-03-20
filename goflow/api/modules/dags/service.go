package dags

import (
	"fmt"
	"log"

	internal_database "github.com/MurilloSSJ/goflow/api/handlers/database"
	DAGModel "github.com/gotoflow/core/entities/dag"
	"github.com/gotoflow/core/handlers/database"
)

type DagService struct {
	database database.Database
}

func ServiceConstructor() *DagService {
	return &DagService{
		database: internal_database.GetDatabase(),
	}
}

func (ds DagService) CreateDag (dagModel DagDTO) {
	dagInsert := &DAGModel.DAG{
		DagId: dagModel.DagId,
	}
	err := ds.database.Insert(dagInsert)
	fmt.Println(err)
	if err != nil {
		log.Fatal(err)
	}
}

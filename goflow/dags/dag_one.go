package main

import (
	DAGModel "lessflow/v2/core/entities/dag"
)

// Função exportada para ser acessada de fora
func GetDAG() *DAGModel.DAG {
	dagParams := &DAGModel.DagParams{
	}
	dag_one := DAGModel.Constructor(*dagParams)
	return dag_one
}
func main() {}
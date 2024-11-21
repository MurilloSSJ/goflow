package main

import (
	DAGModel "lessflow/v2/core/entities/dag"
)

// Função exportada para ser acessada de fora
func GetDAG() *DAGModel.DAG {
	dagParams := &DAGModel.DagParams{
	}
	dag_one := DAGModel.Constructor(*dagParams)
	first_step := func () error {
		return nil
	}
	second_step := func () error {
		return nil
	}
	dag_one.AddNode(first_step)
	dag_one.AddNode(second_step, ["first_step"])
	return dag_one
}
func main() {}
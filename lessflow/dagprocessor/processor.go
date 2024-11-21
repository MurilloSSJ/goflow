package dagprocessor

import (
	"fmt"
	dag_entity "lessflow/v2/core/entities/dag"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"plugin"
)



type DagProcessor struct {
	FilePath string
}

func Constructor(filePath string) *DagProcessor {
	return &DagProcessor{
		FilePath: filePath,
	}
}

func compileDAG(filePath string) (string, error) {
	pluginPath := filePath[:len(filePath)-len(filepath.Ext(filePath))] + ".so"
	fmt.Println(pluginPath)
	cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", pluginPath, filePath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("erro ao compilar DAG: %s\nSaída: %s", err, string(output))
	}
	return pluginPath, nil
}

// loadDAGPlugin carrega o plugin e retorna a estrutura DAG.
func loadDAGPlugin(pluginPath string) (*dag_entity.DAG, error) {
	// Carrega o plugin
	p, err := plugin.Open(pluginPath)
	if err != nil {
		return nil, fmt.Errorf("erro ao carregar plugin: %w", err)
	}

	// Busca a função GetDAG no plugin
	symbol, err := p.Lookup("GetDAG")
	if err != nil {
		return nil, fmt.Errorf("função GetDAG não encontrada no plugin: %w", err)
	}

	// Converte o símbolo para a assinatura esperada
	getDAGFunc, ok := symbol.(func() *dag_entity.DAG)
	if !ok {
		return nil, fmt.Errorf("GetDAG não tem a assinatura esperada")
	}

	// Invoca a função e retorna a DAG
	return getDAGFunc(), nil
}

// ProcessDAG compila e processa uma DAG, retornando seus metadados.
func (dagProcessor *DagProcessor) processDAG() (*dag_entity.DAG, error) {
	// Compila o arquivo em um plugin
	pluginPath, err := compileDAG(dagProcessor.FilePath)
	if err != nil {
		log.Printf("Erro ao compilar DAG (%s): %v\n", dagProcessor.FilePath, err)
		return nil, err
	}
	defer os.Remove(pluginPath) // Limpa o plugin compilado após o uso

	// Carrega o plugin e recupera a DAG
	dag, err := loadDAGPlugin(pluginPath)
	if err != nil {
		log.Printf("Erro ao carregar DAG (%s): %v\n", dagProcessor.FilePath, err)
		return nil, err
	}
	fmt.Println(dag.Nodes)
	return dag, nil
}

func (dagProcessor *DagProcessor) UpdateDag() {
	dag, err := dagProcessor.processDAG()
	if err != nil {
		log.Printf("Erro ao processar DAG (%s): %v\n", dagProcessor.FilePath, err)
		return
	}
	fmt.Println(dag.DagId)
}

func RemoveDag() {
	// Remove DAG
}
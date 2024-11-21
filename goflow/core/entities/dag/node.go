package DAGModel

import (
	"reflect"
	"runtime"
)

type Node struct {
	Name         string
	Dependencies []*Node
	Action       func() error
}

func GetNodeName(i interface{}) string {
    return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
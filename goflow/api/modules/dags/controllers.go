package dags

type DagController struct {
	service *DagService
}

func ControllerConstructor () *DagController {
	return &DagController{
		service: ServiceConstructor(),
	}
}

func (dc *DagController) GetDags() {
	
}

func (dc *DagController) GetDag() {
	
}

func (dc *DagController) CreateDag() {

}

func (dc *DagController) UpdateDag() {
	
}


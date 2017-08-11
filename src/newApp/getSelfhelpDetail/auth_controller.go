package getSelfhelpDetail

import (
	"fmt"
	//"getSelfhelpDetail/exectime"
	workflow "github.com/jabong/floRest/src/common/orchestrator"
)

type AuthController struct {
	id string
}

func (n AuthController) Name() string {
	return "Auth Controller"
}

func (n *AuthController) SetID(id string) {
	n.id = id
}

func (n AuthController) GetID() (id string, err error) {
	return n.id, nil
}

func (n AuthController) Execute(io workflow.WorkFlowData) (workflow.WorkFlowData, error) {
	fmt.Println(io)
	//Start(io)

	//dataDogAgent := monitor.GetInstance()
	//dataDogAgent.Count("_get_Category_hits", 1, []string{"total"}, 1)

	//	io.ExecContext.SetDebugMsg("Category Details - Auth Controller starts", "true")

	//	rc, _ := io.ExecContext.Get(constants.REQUEST_CONTEXT)
	//	logger.Info(fmt.Sprintln("Entered ", n.Name()), rc)
	//
	//	logger.Info(fmt.Sprintln("exiting ", n.Name()), rc)
	//	io.ExecContext.SetDebugMsg("Category Details - Auth Controller Ends", "true")

	return io, nil
}

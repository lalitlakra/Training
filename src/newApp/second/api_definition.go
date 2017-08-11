package second

import (
	"fmt"
	"github.com/jabong/floRest/src/common/constants"
	"github.com/jabong/floRest/src/common/logger"
	"github.com/jabong/floRest/src/common/orchestrator"
	"github.com/jabong/floRest/src/common/ratelimiter"
	"github.com/jabong/floRest/src/common/utils/healthcheck"
	"github.com/jabong/floRest/src/common/versionmanager"
)

type Second struct {
}

func (a *Second) GetVersion() versionmanager.Version {
	return versionmanager.Version{
		Resource: "SECOND",
		Version:  "V1",
		Action:   "GET",
		BucketId: constants.ORCHESTRATOR_BUCKET_DEFAULT_VALUE,
	}
}

func (a *Second) GetOrchestrator() orchestrator.Orchestrator {

	logger.Info("SecondApi BEGINS")
	servOrchestrator := new(orchestrator.Orchestrator)
	servWorkflow := new(orchestrator.WorkFlowDefinition)
	servWorkflow.Create()

	//Create and add execution node AuthController
	authController := new(AuthController)
	authController.SetID("Node 1 : Auth Node")
	authErr := servWorkflow.AddExecutionNode(authController)
	if authErr != nil {
		logger.Error(fmt.Sprintln(authErr))
	}

	GetCategoryDetailsResource := new(GetCategoryDetails)
	GetCategoryDetailsResource.SetID("Node 2 : Get Category Details Node")
	validationErr := servWorkflow.AddExecutionNode(GetCategoryDetailsResource)
	if validationErr != nil {
		logger.Error(fmt.Sprintln(validationErr))
	}

	c1err := servWorkflow.AddConnection(authController, GetCategoryDetailsResource)
	if c1err != nil {
		logger.Error(fmt.Sprintln(c1err))
	}

	servWorkflow.SetStartNode(authController)

	// //Assign the workflow definition to the Orchestrator
	servOrchestrator.Create(servWorkflow)

	logger.Info(servOrchestrator.String())
	logger.Info("Second Pipeline Created")
	return *servOrchestrator
}

func (a *Second) GetHealthCheck() healthcheck.HealthCheckInterface {
	return new(CategoryDetailsHealthCheck)
}

func (a *Second) Init() {

}

func (a *Second) GetRateLimiter() ratelimiter.RateLimiter {
	return nil
}

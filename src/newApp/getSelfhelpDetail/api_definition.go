package getSelfhelpDetail

import (
	"fmt"
	"github.com/jabong/floRest/src/common/constants"
	"github.com/jabong/floRest/src/common/logger"
	"github.com/jabong/floRest/src/common/orchestrator"
	"github.com/jabong/floRest/src/common/ratelimiter"
	"github.com/jabong/floRest/src/common/utils/healthcheck"
	"github.com/jabong/floRest/src/common/versionmanager"
)

type CategoryDetails struct {
}

func (a *CategoryDetails) GetVersion() versionmanager.Version {
	return versionmanager.Version{
		Resource: "CATEGORY",
		Version:  "V1",
		Action:   "GET",
		BucketId: constants.ORCHESTRATOR_BUCKET_DEFAULT_VALUE,
	}
}

func (a *CategoryDetails) GetOrchestrator() orchestrator.Orchestrator {

	logger.Info("CategoryDetailsApi BEGINS")
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
	logger.Info("CategoryDetails Pipeline Created")
	return *servOrchestrator
}

func (a *CategoryDetails) GetHealthCheck() healthcheck.HealthCheckInterface {
	return new(CategoryDetailsHealthCheck)
}

func (a *CategoryDetails) Init() {

}

func (a *CategoryDetails) GetRateLimiter() ratelimiter.RateLimiter {
	return nil
}
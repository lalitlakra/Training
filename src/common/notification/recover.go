package notification

import (
	"fmt"
	//"github.com/jabong/floRest/src/common/utils/logger"
)

func recoverHandler(handler string) {
	if rec := recover(); rec != nil {

		fmt.Println("Error in recover.go")
		//	logger.Error(fmt.Sprintf("[PANIC] occured with %s", handler))
		//	logger.Error(fmt.Sprintf("Reason for panic: %s", rec))
	}
}

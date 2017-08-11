package third

import (
	//"encoding/json"
	"github.com/jabong/floRest/src/common/orchestrator"
	//http "github.com/jabong/floRest/src/common/utils/http"
	//	"github.com/jabong/floRest/src/common/constants"
	//"log"
	//"net/http"
	//	dbconn "common/dao/mysqlconn"
	// //"database/sql"
	"fmt"
	"math/rand"
	"time"

	//"github.com/jabong/floRest/src/common/constants"
	//	"github.com/jabong/floRest/src/common/logger"
	//	"github.com/jabong/floRest/src/common/monitor"
	//"github.com/jabong/floRest/src/common/orchestrator"
	//	"github.com/jabong/floRest/src/common/profiler"
	// //"github.com/go-sql-driver/mysql"
	//	"log"
	//	"reflect"
	//"common/notification"
	//"common/notification/datadog"
	//	utilhttp "github.com/jabong/floRest/src/common/utils/http"
	model "newApp/model/sqldb"
	//"newApp/model/sqldb/Queueinredis"
	//	"strconv"
)

type GetCategoryDetails struct {
	id string
}

// 	{Name: "lalit", Lname: "lakra", EmpId: 1},
// }

func (u *GetCategoryDetails) SetID(id string) {
	u.id = id

}

func (u GetCategoryDetails) GetID() (id string, err error) {
	return u.id, nil
}

func (u GetCategoryDetails) Name() string {
	return "GetCategoryDetails"
}

func (u GetCategoryDetails) Execute(io orchestrator.WorkFlowData) (orchestrator.WorkFlowData, error) {

	//To get values from url
	time.Sleep(time.Millisecond * time.Duration((rand.Intn(30-20) + 20)))
	fmt.Println((rand.Intn(30-20) + 20))

	tr, str := model.IncrementKey("counter")

	fmt.Println(tr)
	fmt.Println(str)

	// rp, _ := io.IOData.Get(constants.REQUEST)
	// appHttpReq, pOk := rp.(*utilhttp.Request)
	// httpReq := appHttpReq.OriginalRequest

	// dataDogAgent := monitor.GetInstance()
	// dataDogAgent.Count("_get_lalit_count", 1, []string{"total"}, 1)

	// p := profiler.NewProfiler()
	//	profiler.StartProfile(p, "nemo_my_category")
	// defer func() {
	// 	t := profiler.EndProfile(p, "nemo_my_category")
	// 	dataDogAgent.Histogram("_get_category_time", float64(t), []string{"getCategoryProfiling"}, 1)
	// }()

	// if !pOk || appHttpReq == nil {
	// 	dataDogAgent.Count("_get_category_info", 1, []string{"Invalid request"}, 1)
	// 	logger.Error("Err: Invalid Request")
	// 	return io, &constants.AppError{Code: constants.IncorrectDataErrorCode, Message: "Err: Invalid Request"}
	// }
	//fmt.Println(httpReq.URL)
	//fmt.Println(httpReq.)
	//	empid := "1"
	//	empid := httpReq.FormValue("id")
	// //conf := new(model.Config)

	// a, _ := model.GetCounter("counter")
	// if a == 0 {
	// 	notification.SendNotification(
	// 		"My app notification ignore",
	// 		fmt.Sprintf("Queue:%s, Count:%s", "myapp", 5),
	// 		[]string{"running", "test"},
	// 		datadog.INFO,
	// 	)
	// }
	// if len(empid) == 0 {
	// 	fmt.Println("zero")
	// }
	//fmt.Println("here")
	//	var empId int
	//	fmt.Println(empid)

	//	empId, _ = strconv.Atoi(empid)

	//fmt.Println(empId)
	//empId := 1
	//ret SqlDbInterface
	//var err *SqlDbError

	//ret, _ := model.Get(conf)
	//args := ""

	//err = rows.Err()
	//if err != nil {
	///	log.Fatal(err)
	//}

	// //x := []ob

	// //io.IOData.Set("RESULT", obj)
	// //io.IOData.Set("RESULT", "lalit")
	// fmt.Println("3")
	io.IOData.Set("RESULT", "2")
	return io, nil
}

// to prepare get list of active categories response

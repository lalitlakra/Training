package getSelfhelpDetail

import (
	"encoding/json"
	"github.com/jabong/floRest/src/common/orchestrator"
	//http "github.com/jabong/floRest/src/common/utils/http"
	//	"github.com/jabong/floRest/src/common/constants"
	//"log"
	"math/rand"
	"net/http"
	"sync"
	//	"reflect"
	//	dbconn "common/dao/mysqlconn"
	// //"database/sql"
	"fmt"
	"io/ioutil"
	"time"
	//"github.com/jabong/floRest/src/common/constants"
	//	"github.com/jabong/floRest/src/common/logger"
	"github.com/jabong/floRest/src/common/monitor"
	//"github.com/jabong/floRest/src/common/orchestrator"
	"github.com/jabong/floRest/src/common/profiler"
	// //"github.com/go-sql-driver/mysql"
	//	"log"
	//	"reflect"
	//"common/notification"
	//"common/notification/datadog"
	//utilhttp "github.com/jabong/floRest/src/common/utils/http"
	//	"newApp/model/sqldb"
	model "newApp/model/sqldb"
	"strconv"
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

func ex(channel chan *http.Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * time.Duration((rand.Intn(10) + 20)))
	resp, err := http.Get("http://localhost:9090/newApp/v1/second")
	fmt.Println(err)
	channel <- resp
	wg.Done()
}
func ex1(channel chan *http.Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * time.Duration((rand.Intn(10) + 20)))
	resp, err := http.Get("http://localhost:9090/newApp/v1/third")
	fmt.Println(err)
	channel <- resp
	wg.Done()
}
func (u GetCategoryDetails) Execute(io orchestrator.WorkFlowData) (orchestrator.WorkFlowData, error) {

	//To get values from url

	tr, str := model.IncrementKey("counter")

	fmt.Println(tr)
	fmt.Println(str)

	//To get data from url of api request
	//	rp, _ := io.IOData.Get(constants.REQUEST)
	//	appHttpReq, pOk := rp.(*utilhttp.Request)
	//	httpReq := appHttpReq.OriginalRequest

	//TO SET DATADAG GRAPH
	dataDogAgent := monitor.GetInstance()
	dataDogAgent.Count("_get_lalit_count", 1, []string{"total"}, 1)

	p := profiler.NewProfiler()
	profiler.StartProfile(p, "nemo_my_category")
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

	//TO GET ID FIELD VALUE FROM URL
	//empid := httpReq.FormValue("id")
	// //conf := new(model.Config)

	//HIT ANOTHER API AND GET RESPONSE IN RESP
	var channel = make(chan *http.Response)
	var wg sync.WaitGroup
	//go ex()
	//var resp *http.Response
	//var resp1 *http.Response
	wg.Add(1)

	go ex(channel, &wg)

	wg.Add(1)

	go ex1(channel, &wg)
	resp := <-channel
	resp1 := <-channel
	wg.Wait()
	//var err error
	//resp, err = http.Get("http://localhost:9090/newApp/v1/second/?id=4")

	//resp1, _ := http.Get("http://localhost:9090/newApp/v1/third/?id=2")

	fmt.Println("Data", resp)
	//fmt.Println(err)
	//ALWAYS CLOSE RESP.BODY
	defer resp.Body.Close()
	defer resp1.Body.Close()
	//TO GET BODY OF RESPONSE
	body, _ := ioutil.ReadAll(resp.Body)
	body1, _ := ioutil.ReadAll(resp1.Body)
	//TO GET DATA FROM RESPONSE
	son := []byte(body)
	son2 := []byte(body1)
	var datamap model.ResponseStruct
	var datamap1 model.ResponseStruct

	fmt.Println("get:\n", string(body))
	fmt.Println("get:\n", string(body1))
	//data := sqldb.Test{}

	err1 := json.Unmarshal(son, &datamap)
	err2 := json.Unmarshal(son2, &datamap1)
	if err1 != nil {
		fmt.Println(err1)

	}
	if err2 != nil {
		fmt.Println(err1)

	}
	//var data = model.Test{}
	// data := make(map[string]interface{})
	//data = datamap["data"].(model.Test)

	//fmt.Println(data.Ename)

	//fmt.Println(data["id"])
	//	var obj []model.Test
	//	var ob1 = model.Test{Ename: datamap.Data[0].Ename, EmpId: datamap.Data[0].EmpId}
	//	obj = append(obj, ob1)
	a, _ := strconv.Atoi(datamap.Data)
	b, _ := strconv.Atoi(datamap1.Data)
	var ob1 = model.Test{Ename: "XYZ", EmpId: a + b}

	//	var ob2 = model.Test{Ename: datamap1.Data[0].Ename, EmpId: datamap1.Data[0].EmpId}
	//	obj = append(obj, ob2)
	fmt.Println("DATA", datamap.Data)
	fmt.Println("DATA1", datamap1.Data)
	//fmt.Println(data.EmpId)
	//fmt.Println(data.Ename)

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
	//var obj model.Test

	//obj, ok := datamap["data"].(string)
	//obj = append(obj, model.Test(datamap1["data"]))
	//result := datamap["data"].(model.Test)
	//fmt.Println(result)
	// //io.IOData.Set("RESULT", obj)
	// //io.IOData.Set("RESULT", "lalit")
	// fmt.Println("3")
	//fmt.Println(ok)
	//fmt.Println(reflect.TypeOf(datamap["data"]))
	io.IOData.Set("RESULT", ob1)
	//io.IOData.Set("RESULT", datamap1["data"])

	return io, nil
}

// to prepare get list of active categories response

package sqldb

import (
	dbconn "common/dao/mysqlconn"
	// //"database/sql"
	//	"database/sql"
	"fmt"
	//	_ "github.com/go-sql-driver/mysql"

	// //"github.com/go-sql-driver/mysql"
	//	"log"
	//  "newApp/getSelfhelpDetail"
	//
	"bytes"
	"database/sql"
	"github.com/go-redis/redis"
	"strconv"
)

type Test struct {
	Ename string `json:"name" bson:"name"`
	//Lname string `json:"lname" bson:"lname"`
	EmpId int `json:"id" bson:"id"`
}
type Test1 struct {
	Ename string `json:"name" bson:"name"`
	//Lname string `json:"lname" bson:"lname"`
	//EmpId  int `json:"id" bson:"id"`
	//Salary int `json:"salary" `
}

type ResponseStruct struct {
	Status interface{}
	Data   string //[]Test
	Other  interface{}
}

//type Employees struct {

//}

func NewClient() redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return *client
}
func Example() {

	client := NewClient()

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	err = client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	m := make(map[string]interface{})
	m["id"] = 1
	m["name"] = "lalit"
	client.HMSet("emp:1", m)
	fmt.Println(client.HGetAll("emp:1"))
	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exitst")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}

}

func GetfromRedis(id int, query string) (bool, []Test) {

	var obj []Test
	var buffer bytes.Buffer
	//client := NewClient()

	buffer.WriteString("emp:")
	buffer.WriteString(strconv.Itoa(id))
	//fmt.Println("123 ")
	m, _ := dbconn.Client.HGetAll(buffer.String()).Result()
	//m, _ := client.HGetAll(buffer.String()).Result()

	//fmt.Println(m)
	//fmt.Println(buffer.String())
	err := -1
	if len(m) == 0 {
		err = 1
	} else {
		err = 0
	}

	if err == 1 {
		rows, err := dbconn.Mysqldb.Query(query, id)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		return false, SetinRedis(id, rows)
		//return false, obj
	} else {
		//fmt.Println("about to")
		temp, _ := strconv.Atoi(m["id"])
		o1 := Test{EmpId: temp, Ename: m["name"]}
		obj = append(obj, o1)
		return true, obj

	}

}

func SetinRedis(ix int, rows *sql.Rows) []Test {
	var obj []Test

	var (
		id   int
		name string
	)

	//client := NewClient()
	m := make(map[string]interface{})

	for rows.Next() {

		var buffer bytes.Buffer
		//client := NewClient()

		buffer.WriteString("emp:")
		buffer.WriteString(strconv.Itoa(ix))
		///	fmt.Println(buffer.String())
		err := rows.Scan(&name, &id)

		m["id"] = strconv.Itoa(id)
		m["name"] = name
		dbconn.Client.HMSet(buffer.String(), m)
		//temp, _ := strconv.Atoi(m["id"])
		o1 := Test{EmpId: id, Ename: m["name"].(string)}
		obj = append(obj, o1)

		//fmt.Println(client.HGetAll(buffer.String()).Result())
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		///var Arr *Employees

		//Employees.obj = append(Employees.obj, o1)

	}
	return obj

}

func GetEmployee(empId int) []Test {

	//	Example()
	query := `SELECT   EmpName ,EmpId from Employee where EmpId=?`
	//dbconn.Initialize()
	//fmt.Println(query)
	//query := ' Select * from Employee '
	//stmt, err := dbconn.Mysqldb.Prepare(query)

	//	stmt,_:= dbconn.Mysqldb.Prepa
	// u, err := url.Parse(s)
	// if err != nil {
	// 	panic(err)
	// }
	//query1 := `use MyDatabase  `
	//var query1 string
	//query1 =  use MyDatabase '
	//_, err1 := dbconn.Mysqldb.Query(query1)

	//fmt.Println("db used", err1)
	_, ob := GetfromRedis(empId, query)

	//fmt.Println("after query")

	//InsertinQueue()
	//fmt.Println("Queue data")
	//GetfromQueue()
	//res, err := stmt.Exec(empname, empid)
	// //db, err := sql.Open("mysql",
	// //"root:jabong@123@tcp(localhost:3306)/MyDatabase")
	// fmt.Println("1")

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// //defer db.Close()

	// //rows, err := db.Query("select EmpName, EmpId from Employee ")

	// defer rows.Close()

	return ob

}

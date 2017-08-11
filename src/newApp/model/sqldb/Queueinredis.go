package sqldb

import (
	dbconn "common/dao/mysqlconn"
	"fmt"

	"encoding/json"
	"github.com/go-redis/redis"
	"strconv"
)

//var maxValue string
//maxValue = "maxValue"
//returns value of counter
func GetCounter(key string) (int64, error) {
	//client := NewClient()
	return dbconn.Client.Get(key).Int64()
}

//returns max limit of counter
func GetMaxValue() (int64, error) {
	//client := NewClient()
	return dbconn.Client.Get("maxValue").Int64()

}

//set max limit of counter
func SetMaxValue() {
	//client := NewClient()
	fmt.Println("hello max")
	//client.Set(key, value, expiration)
	dbconn.Client.Set("maxValue", 100, 0)
}

//Set initial value and key of counter

func SetCounter(key string, value int) {
	//client := NewClient()
	dbconn.Client.Set(key, value, 0)

}

//increments the counter
func IncrementKey(key string) (bool, string) {
	//client := NewClient()
	var ans int
	ans = -1

	err := dbconn.Client.Watch(func(tx *redis.Tx) error {

		n, err := tx.Get(key).Int64()
		if err != nil && err != redis.Nil {
			return err
		}
		//fmt.Println("here 1")
		b, _ := GetMaxValue()
		fmt.Println("value of counter ", n)

		if n < b {
			_, err = tx.Pipelined(func(pipe redis.Pipeliner) error {
				pipe.Set(key, strconv.FormatInt(n+1, 10), 0)
				ans = 1
				//	fmt.Println("here 2")
				return nil
			})

		} else {
			//fmt.Println("here 3")
			ans = 0

		}
		return err
	}, key)

	if err == redis.TxFailedErr {
		//	return incr(key)
	}

	//return err

	//a, _ := GetCounter(key)
	//client.Wait(1, 10)
	// client.Watch(fn, ...)
	// client.

	//client.multi()
	//if a < b {
	//client.Incr(key)
	//	return true, "Incremented"

	//}
	//client.Command("exec")
	if ans == 1 {
		return true, "Incremented"
	} else if ans == 0 {
		return false, "Value greater than max"
	} else {
		return false, "No operation"
	}

}

func InsertinQueue() {

	//client := NewClient()
	test1 := Test{Ename: "Rishab", EmpId: 1}
	test2 := Test{Ename: "Rishab Kr", EmpId: 2}
	//client.LPush("list", test1)
	//	m := make(map[string]interface{})
	a, _ := json.Marshal(test1)
	b, _ := json.Marshal(test2)
	//client.HMGet(key, ...)

	//	m["id"] = 123
	//	m["name"] = "Lalit"

	dbconn.Client.LPush("list", a)
	dbconn.Client.LPush("list", b)
	//	fmt.Println("hello")
	//client.LPush("list", "lakra")

}

func GetfromQueue() {
	//client := NewClient()
	//i, _ := client.LLen("list").Result()
	//var j int64
	//for j = 0; j < i; j++ {
	//	fmt.Println(client.LPop("list"))
	//}

	//fmt.Println("dssd")
	fmt.Println(dbconn.Client.LRange("list", 0, 1))
	//fmt.Println(client.LPop("list"))
}

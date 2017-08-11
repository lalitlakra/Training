package mysqlconn

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jabong/floRest/src/common/sqldb"
	"os"
)

var Mysqldb sqldb.SqlDbInterface
var MysqlSlavedb sqldb.SqlDbInterface
var Client redis.Client

func Initialize() {

	fmt.Println("created")
	conErr := new(sqldb.SqlDbError)

	dbInfo := new(sqldb.Config)
	dbInfo.DriverName = "mysql"
	dbInfo.Username = os.Getenv("MYSQL_USER_NAME")
	dbInfo.Password = os.Getenv("MYSQL_PASSWORD")
	dbInfo.Host = os.Getenv("MYSQL_MASTER_HOST")
	dbInfo.Port = os.Getenv("MYSQL_MASTER_PORT")
	dbInfo.Dbname = os.Getenv("MYSQL_DBNAME")
	dbInfo.Timezone = "Local"
	dbInfo.MaxOpenCon = 10
	dbInfo.MaxIdleCon = 5
	Mysqldb, _ = sqldb.Get(dbInfo)

	if conErr.ErrCode != "" {
		panic("Mysql Connection Error.")
	}

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	//	fmt.Println(client)
	Client = *client

	dbInfo = new(sqldb.Config)
	dbInfo.DriverName = "mysql"
	dbInfo.Username = os.Getenv("MYSQL_SLAVE_USER_NAME")
	dbInfo.Password = os.Getenv("MYSQL_SLAVE_PASSWORD")
	dbInfo.Host = os.Getenv("MYSQL_SLAVE_HOST")
	dbInfo.Port = os.Getenv("MYSQL_SLAVE_PORT")
	dbInfo.Dbname = os.Getenv("MYSQL_DBNAME")
	dbInfo.Timezone = "Local"
	dbInfo.MaxOpenCon = 10
	dbInfo.MaxIdleCon = 5
	MysqlSlavedb, _ = sqldb.Get(dbInfo)

	if conErr.ErrCode != "" {
		panic("Mysql Slave Connection Error.")
	}
}

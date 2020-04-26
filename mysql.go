// gomysql driver
package mylib

import (
	"fmt"
	"log"
	"time"
)
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

var db = &sql.DB{}

// @doc 连接数据库
// @param dataSourceName::root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4
func Connect(driverName string, dataSourceName string) {
	var err error
	db, err = sql.Open(driverName, dataSourceName)
	fmt.Printf("err=%v\n", err)
	fmt.Printf("err=%v\n", db)
	QueryTest()
}

func QueryTest() {
	start := time.Now()
	rows, _ := db.Query("SELECT uid,username FROM USER")
	defer rows.Close()
	for rows.Next() {
		var name string
		var id int
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("name:%s ,id:is %d\n", name, id)
	}
	end := time.Now()
	fmt.Println("方式1 query total time:", end.Sub(start).Seconds())
}

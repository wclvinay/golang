package db
import (
     "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "fmt"
)
var Connection *sql.DB
func init() {
	Connection,_ = sql.Open("mysql", "root:narender123@/customer_service")
	err := Connection.Ping()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("db is connected")
	}
}	
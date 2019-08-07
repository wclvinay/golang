package db
import (
     "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "fmt"
)
func CreateCon() *sql.DB {
	db, err := sql.Open("mysql", "root:narender123@/customer_service")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("db is connected")
	}
	//defer db.Close()
	// make sure connection is available
	err = db.Ping()
	fmt.Println(err)
	if err != nil {
		fmt.Println("MySQL db is not connected")
		fmt.Println(err.Error())
	}
	return db
}
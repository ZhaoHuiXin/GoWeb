package main
import(

	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

func main() {
	db, err := sql.Open("mysql", "root:666666@tcp" +
		"(localhost:3306)/test?charset=utf8")
	checkErr(err)
	// insert data
	//stmt, err := db.Prepare("insert into userinfo(username,department,created) values(?,?,?)")
	//checkErr(err)
	//res, err := stmt.Exec("zhaohuixin","golang","2018-04-25")
	//checkErr(err)
	//fmt.Println(res)

	//update data
	//stmt, err := db.Prepare("update userinfo set username=? where uid=?")
	//checkErr(err)
	//stmt.Exec("lucy",1)
	//print(res)

	//query data
	rows, err := db.Query("select * from userinfo")
	fmt.Println(rows)
	checkErr(err)
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
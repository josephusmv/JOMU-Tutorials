package main

import (
	"database/sql"
	"flag"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

const cStrDfaultIP = "127.0.0.1"
const cStrDfaultPort = "3306"

const cStrDefaultDBName = "DB_Player"
const cStrDefaultTBLName = "TBL_PLAYERS"

//CREATE USER jomu@localhost IDENTIFIED BY '123456';
//GRANT ALL on DB_Player.* to jomu@localhost;
//CREATE DATABASE IF NOT EXISTS DB_Player DEFAULT CHARSET utf8 COLLATE utf8_general_ci;
func connectDB(dbname, userpwd string) *sql.DB {
	connStr := fmt.Sprintf("%s@tcp(%s:%s)/%s?charset=utf8", userpwd, cStrDfaultIP, cStrDfaultPort, dbname)

	fmt.Println(connStr)
	db, err := sql.Open("mysql", connStr)

	checkErr(err)
	if db == nil {
		panic("Empty DB, failed.")
	}

	err = db.Ping()
	checkErr(err)

	return db
}

func main() {

	add := flag.Bool("add", false, "Add New Values.")
	qry := flag.Bool("qry", false, "Query New Values.")
	del := flag.Bool("del", false, "Delete New Values.")
	test := flag.Bool("test", false, "use test values.")
	flag.Parse()
	if *add {
		doAdd(*test)
	}

	if *del {
		doDelete(*test)
	}

	if *qry {
		doSelect(*test)
	}
}

type MyslDBFieldValue struct {
	fldType  string
	fldName  string
	fldValue interface{}
}

func doAdd(test bool) {
	var inputs map[string]*MyslDBFieldValue
	if test {
		inputs = getInputsAuto()
	} else {
		inputs = getInputs()
	}

	db := connectDB(cStrDefaultDBName, "jomu:123456")
	defer db.Close()

	doInsert(db, cStrDefaultTBLName, inputs)
}
func doSelect(test bool) {

	db := connectDB(cStrDefaultDBName, "jomu:123456")
	defer db.Close()

	stmtStr := fmt.Sprintf("SELECT * FROM %s where player_name like ?", cStrDefaultTBLName)
	fmt.Println(stmtStr)

	stmt, errprp := db.Prepare(stmtStr)
	checkErr(errprp)

	fldValue := "%it"
	rows, errq := stmt.Query(fldValue)
	checkErr(errq)
	defer rows.Close()

	var id int
	var name string
	var online string
	var addr string
	rslts := make([]interface{}, 4)
	rslts[0] = &id
	rslts[1] = &name
	rslts[2] = &online
	rslts[3] = &addr

	for rows.Next() {
		err := rows.Scan(rslts...)
		checkErr(err)
	}

	fmt.Printf("%d, %s, %s, %s\n", id, name, online, addr)
}

type MySqlExeRslt struct {
	RowsAffected int64
	LastInsertId int64
}

func getInputsAuto() map[string]*MyslDBFieldValue {
	inputs := make(map[string]*MyslDBFieldValue)

	fieldID := MyslDBFieldValue{
		fldName:  "id",
		fldType:  "int(11)",
		fldValue: "100023",
	}
	inputs["id"] = &fieldID

	fieldName := MyslDBFieldValue{
		fldName:  "player_name",
		fldType:  "varchar(255)",
		fldValue: "john smith",
	}
	inputs["player_name"] = &fieldName

	fieldOnline := MyslDBFieldValue{
		fldName:  "online",
		fldType:  "char(1)",
		fldValue: "Y",
	}
	inputs["online"] = &fieldOnline

	fieldAddress := MyslDBFieldValue{
		fldName:  "address",
		fldType:  "varchar(255)",
		fldValue: "34# TopStreet baker land, waterloo",
	}
	inputs["address"] = &fieldAddress

	return inputs
}

func getInputs() map[string]*MyslDBFieldValue {
	inputs := make(map[string]*MyslDBFieldValue)

	fmt.Print("\nID: ")
	var id int
	fmt.Scanln(&id)
	fieldID := MyslDBFieldValue{
		fldName:  "id",
		fldType:  "int(11)",
		fldValue: id,
	}
	inputs["id"] = &fieldID

	fmt.Print("\nName: ")
	var name string
	fmt.Scanln(&name)
	fieldName := MyslDBFieldValue{
		fldName:  "player_name",
		fldType:  "varchar(255)",
		fldValue: name,
	}
	inputs["player_name"] = &fieldName

	fmt.Print("\nOnline: ")
	var online string
	fmt.Scanln(&online)
	fieldOnline := MyslDBFieldValue{
		fldName:  "online",
		fldType:  "char(1)",
		fldValue: online,
	}
	inputs["online"] = &fieldOnline

	fmt.Print("\nAddress: ")
	var addr string
	fmt.Scanln(&addr)
	fieldAddress := MyslDBFieldValue{
		fldName:  "address",
		fldType:  "varchar(255)",
		fldValue: addr,
	}
	inputs["address"] = &fieldAddress

	fmt.Printf("\n %d, %s, %s, %s\n", id, name, online, addr)

	return inputs
}

func doInsert(db *sql.DB, dbname string, inputs map[string]*MyslDBFieldValue) {
	stmtStr := fmt.Sprintf("INSERT INTO %s (", dbname)

	valueStr := " VALUE( "
	var count int
	values := make([]interface{}, len(inputs))
	for k := range inputs {
		stmtStr += k
		valueStr += "?"
		if count < len(inputs)-1 {
			stmtStr += ", "
			valueStr += ", "
		}
		values[count] = inputs[k].fldValue
		count++
	}

	stmtStr += ")" + valueStr + ")"

	fmt.Println(stmtStr)

	stmt, errprp := db.Prepare(stmtStr)
	checkErr(errprp)

	rslt, errExe := stmt.Exec(values...)
	checkErr(errExe)

	result := MySqlExeRslt{}
	result.LastInsertId, _ = rslt.LastInsertId()
	result.RowsAffected, _ = rslt.RowsAffected()

	fmt.Printf("\nInsert Result: %v \n", result)
}

func doDelete(test bool) {
	var inputs *MyslDBFieldValue
	if test {
		inputs = getIndexAuto()
	} else {
		inputs = getIndex()
	}

	db := connectDB(cStrDefaultDBName, "jomu:123456")
	defer db.Close()

	doDeleteByIndexSingle(db, cStrDefaultTBLName, inputs)
}

func getIndexAuto() *MyslDBFieldValue {
	return &MyslDBFieldValue{
		fldName:  "player_name",
		fldType:  "varchar(255)",
		fldValue: "john smith",
	}
}

func getIndex() *MyslDBFieldValue {
	var name string
	fmt.Println("\nEnter Index Name for delete: ")
	fmt.Scanln(&name)
	return &MyslDBFieldValue{
		fldName:  "player_name",
		fldType:  "varchar(255)",
		fldValue: name,
	}
}

func doDeleteByIndexSingle(db *sql.DB, dbname string, indexField *MyslDBFieldValue) {
	stmtStr := fmt.Sprintf("DELETE FROM %s where %s=?", dbname, indexField.fldName)

	fmt.Println(stmtStr)

	stmt, errprp := db.Prepare(stmtStr)
	checkErr(errprp)

	rslt, errExe := stmt.Exec(indexField.fldValue)
	checkErr(errExe)

	result := MySqlExeRslt{}
	result.LastInsertId, _ = rslt.LastInsertId()
	result.RowsAffected, _ = rslt.RowsAffected()

	fmt.Printf("\nInsert Result: %v \n", result)
}

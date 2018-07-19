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

type MysqlDBField struct {
	fldName    string
	fldType    string
	notnull    bool
	index      bool
	primKey    bool
	frgnKey    bool
	frgnConst  string
	unique     bool
	deflt      bool
	defltValue string
	autoinc    bool
}

func formTestColumnData() []MysqlDBField {
	fields := make([]MysqlDBField, 4)
	fields[0].fldName = "id"
	fields[0].fldType = "int"
	fields[0].primKey = true
	fields[0].autoinc = true

	fields[1].fldName = "player_name"
	fields[1].fldType = "varchar(255)"
	fields[1].deflt = true
	fields[1].defltValue = "ananoym player"

	fields[2].fldName = "online"
	fields[2].fldType = "char(1)"
	fields[2].notnull = true
	fields[2].deflt = true
	fields[2].defltValue = "N"

	fields[3].fldName = "address"
	fields[3].fldType = "varchar(255)"

	return fields
}

func createTableIfNotExist(db *sql.DB, tablename string, fields []MysqlDBField) {
	stmtStr := "CREATE TABLE IF NOT EXISTS " + tablename + " ( \n"

	var primaryKey string
	foreignKeys := make(map[string]string)
	for i, field := range fields {
		stmtStr += field.fldName + " " + field.fldType

		if field.notnull {
			stmtStr += " NOT NULL "
		}
		if field.autoinc {
			stmtStr += " AUTO_INCREMENT "
		}

		if field.deflt {
			stmtStr += " DEFAULT"
			if field.defltValue != "" {
				stmtStr += " '" + field.defltValue + "'"
			}
		}

		if field.primKey {
			primaryKey = field.fldName //not consider CONSTRAINT pk_PersonID PRIMARY KEY (Id_P,LastName)
		} else {
			if field.frgnKey {
				foreignKeys[field.fldName] = field.frgnConst
			} else if field.unique {
				stmtStr += "  "
			}
		}

		if i+1 < len(fields) {
			stmtStr += ", \n"
		}
	}

	if len(primaryKey) > 0 {
		stmtStr += ",\nPRIMARY KEY (" + primaryKey + ")"
	}

	for k, v := range foreignKeys {
		stmtStr += ",\n FOREIGN KEY (" + k + ") REFERENCES " + v
	}

	stmtStr += "\n)"

	fmt.Println(stmtStr)

	stmt, err := db.Prepare(stmtStr)
	checkErr(err)

	_, err = stmt.Exec()
	checkErr(err)
}

func dropTableIfExists(db *sql.DB, dbname string) {
	stmt, err := db.Prepare("DROP TABLE IF EXISTS " + dbname)
	checkErr(err)

	_, err = stmt.Exec()
	checkErr(err)
}

func main() {

	drop := flag.Bool("drop", false, "Drop created test table.")
	flag.Parse()

	db := connectDB(cStrDefaultDBName, "jomu:123456")
	fields := formTestColumnData()

	createTableIfNotExist(db, "TBL_PLAYERS", fields)

	if *drop {
		dropTableIfExists(db, "TBL_PLAYERS")
	}

	db.Close()
}

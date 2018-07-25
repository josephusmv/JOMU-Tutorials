package main

import "database/sql"

type DBOperator struct {
	db *sql.DB
}

func NewDBOperator() *DBOperator {
	return &DBOperator{}
}

func (dbo *DBOperator) ConnectDB(dbname string, connStr string) {

}

func (dbo *DBOperator) Close() {
	dbo.db.Close()
}

func (dbo *DBOperator) CreateTable(tableName string, fields []*DBFieldDefn) *DBTableDefn {
	return nil
}

func (dbo *DBOperator) CreateTempTable(tableName string, fields []*DBFieldDefn) *DBTableDefn {
	return nil
}

func (dbo *DBOperator) DropTable(tableName string) *DBTableDefn {
	return nil
}

//DoSelect SELECT cols FROM tableName WHERE  wherStr(bindVars)
//	wherStr - inform or: "COL1=? AND COL2 like '?'", do not includes "WHERE" key words
//	bindVars - all search variables are using variable binding, if you need a fuzzy query, include % in the value, like: "J% Smith"
func (dbo *DBOperator) DoSelect(tableName string, cols []string, wherStr string, bindVars []interface{}) (*DBExeRslt, []DBRowValues) {
	return nil, nil
}

//DoUpdate UPDATE tableName SET cols[i] = bindVars[i] WHERE wherStr(bindVars)
//	wherStr - inform or: "COL1=? AND COL2 like '?'"
//	bindVars - all search variables are using variable binding, if you need a fuzzy query, include % in the value, like: "J% Smith"
//			both set values and where condition values are all expected to be stored in bindVars
func (dbo *DBOperator) DoUpdate(tableName string, cols []string, wherStr string, bindVars []interface{}) (*DBExeRslt, []DBRowValues) {
	return nil, nil
}

//DoInsertInto INSERT INTO tableName (cols) VALUES(bindVars)
//	bindVars - all search variables are using variable binding, if you need a fuzzy query, include % in the value, like: "J% Smith"
func (dbo *DBOperator) DoInsertInto(tableName string, cols []string, bindVars []interface{}) *DBExeRslt {
	return nil
}

//DoDelete DELETE FROM tableName (cols) WHERE wherStr(bindVars)
//	wherStr - inform or: "COL1=? AND COL2 like '?'"
//	bindVars - all search variables are using variable binding, if you need a fuzzy query, include % in the value, like: "J% Smith"
//			both set values and where condition values are all expected to be stored in bindVars
func (dbo *DBOperator) DoDelete(tableName string, wherStr string, bindVars []interface{}) *DBExeRslt {
	return nil
}

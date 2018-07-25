package main

type DBTableDefn struct {
	Name       string
	Index      string
	PrimaryKey string
	ForeignKey string
	FieldTypes []string
}

type DBFieldDefn struct {
	FldName      string
	FldType      string
	NotNull      bool
	IsIndex      bool
	IsPrimKey    bool
	IsFrgnKey    bool
	FrgnTable    string
	IsUnique     bool
	HasDefault   bool
	DefaultValue string
	AutoIncr     bool
}

/*
type DBFieldValue struct {
	fldType  string
	fldName  string
	fldValue interface{}
}
*/

type DBRowValues []DBFieldValue

type DBExeRslt struct {
	RowsAffected int64
	LastInsertId int64
}

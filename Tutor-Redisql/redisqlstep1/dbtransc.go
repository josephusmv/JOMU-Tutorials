package main

//DBTransaction A transation for operations on DB. Users should only connect to one DB instance with one transaction.
//	DB transation maintains all synchronizing works agains one DB.
type DBTransaction struct {
	dbopr *DBOperator
}

func StartDBTransaction() (*DBTransaction, error) {
	return nil, nil
}

func (dbt *DBTransaction) Stop() {
	if dbt.dbopr == nil {
		return
	}

	dbt.dbopr.Close()
}

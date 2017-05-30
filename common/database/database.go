package database

import (
	"fmt"
	"database/sql"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//Primary Database Object
type Db struct {
	Connection *sqlx.DB
	DbConn     string
	DbType     string
}

type Tx struct {
	Transaction *sqlx.Tx
}

var ErrNoRows = sql.ErrNoRows

//Connect to database from config and Ping the connection
func Connect(dbType, dbConn string) (*Db, error) {
	db, err := sqlx.Connect(dbType, dbConn)
	if err != nil {
		return nil, fmt.Errorf("DB open connection error. Error:", err.Error())
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("DB ping connection error. Error:", err.Error())
	}

	database := &Db{
		Connection: db,
		DbConn: dbConn,
		DbType: dbType,
	}

	return database, nil
}

func (d *Db) DoAction() (*sqlx.DB) {
	return d.Connection
}

func (d *Db) StartTransaction() (*Tx, error) {
	tx, err := d.Connection.Beginx()
	if err != nil {
		return nil, fmt.Errorf("Failed to start database transaction. Error: ", err.Error())
	}

	trx := &Tx{
		Transaction: tx,
	}

	return trx, nil
}

func (tx *Tx) Rollback() error {
	return tx.Transaction.Rollback()
}

func (tx *Tx) Commit() error {
	return tx.Transaction.Commit()
}

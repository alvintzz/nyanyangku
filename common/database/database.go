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
}
type Tx struct {
	Transaction *sqlx.Tx
}

var ErrNoRows = sql.ErrNoRows
var dbList = map[string]Db{}

//Connect to database from config and Ping the connection
func Connect(name, dbType, dbConn string) error {
	if _, ok := dbList[name]; ok {
		return fmt.Errorf("Database %s is already initialized.", name)
	}

	db, err := sqlx.Connect(dbType, dbConn)
	if err != nil {
		return fmt.Errorf("DB open connection error. Error:", err.Error())
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("DB ping connection error. Error:", err.Error())
	}

	dbList[name] = Db{
		Connection: db,
	}
	return nil
}

//Set the database connection directly
func Set(name string, conn *sqlx.DB) error {
	dbList[name] = Db{
		Connection: conn,
	}
	return nil
}

func Get(name string) (Db, error) {
	if db, ok := dbList[name]; ok {
		return db, nil
	}
	return Db{}, fmt.Errorf("Engine %s haven't initialized.", name)
}

func (d Db) DoAction() (*sqlx.DB) {
	return d.Connection
}

func (d Db) StartTransaction() (*Tx, error) {
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

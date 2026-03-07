package database

import "database/sql"

type DbManager struct {
	RTV *sql.DB
	Billing *sql.DB
	Auth *sql.DB
}
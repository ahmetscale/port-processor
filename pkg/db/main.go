package db

import (
	"gorm.io/gorm"
)

// DB is a wrapper for gorm.DB
type DB struct {
	*gorm.DB
}

// New creates a new DB
func New(dialect, dsn string, logMode bool) *DB {
	db, err := Open(dsn, dialect)
	if err != nil {
		panic(err)
	}
	if logMode {
		db = db.Debug()
	}
	var conn DB
	conn.DB = db
	return &conn
}

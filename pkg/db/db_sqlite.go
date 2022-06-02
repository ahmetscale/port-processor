//go:build sqlite

// Package db +build sqlite
package db

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func Open(dsn string, adapter string) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(dsn), &gorm.Config{})
}

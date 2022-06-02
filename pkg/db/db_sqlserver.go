//go:build sqlserver

// Package db +build sqlserver
package db

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func Open(dsn string, adapter string) (*gorm.DB, error) {
	return gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
}

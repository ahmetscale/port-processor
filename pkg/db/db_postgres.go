//go:build postgres

// Package db +build postgres
package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Open(dsn string, adapter string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
}

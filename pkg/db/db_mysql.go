//go:build mysql

// Package db +build mysql
package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Open(dsn string, adapter string) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

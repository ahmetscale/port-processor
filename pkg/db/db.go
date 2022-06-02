//go:build !sqlite && !mysql && !postgres && !sqlserver

// Package db Package pkg +build !sqlite,!mysql,!postgres,!sqlserver
package db

import "log"

import (
	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func Open(dsn string, adapter string) (*gorm.DB, error) {
	if adapter == "mysql" {
		return gorm.Open(mysql.Open(dsn), &gorm.Config{})
	} else if adapter == "postgres" {
		return gorm.Open(postgres.Open(dsn), &gorm.Config{})
	} else if adapter == "sqlite" {
		return gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	} else if adapter == "sqlserver" {
		return gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	} else {
		log.Fatal("Unknown adaptor: ", adapter)
	}
	return nil, nil
}

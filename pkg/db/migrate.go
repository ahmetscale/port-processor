package db

func Migrate(db *DB, dst ...interface{}) (err error) {
	return db.AutoMigrate(dst...)
}

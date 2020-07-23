package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // For using Postgres as database
	"github.com/ramezanpour/users/utilities/config"
)

// Db is the global object for accessing database
var Db *gorm.DB

// Init initializes a connection to the database
func Init() {
	connectionString := config.GetConfig().DbConnectionString
	var err error
	Db, err = gorm.Open("postgres", connectionString)

	if err != nil {
		panic(err)
	}

	Db.SingularTable(false)
}

// Close closes all connections to the DB
func Close() {
	Db.Close()
}

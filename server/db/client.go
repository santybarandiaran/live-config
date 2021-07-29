package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	//revive:disable-next-line:blank-imports this is needed by gorm to do the DB connection
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"strconv"
)

// Init executes the setup related with the DB connection
func Init() (*gorm.DB, error) {
	conf := New()
	dbType, dbConnection := getDBConnection(conf)

	db, err := gorm.Open(dbType, dbConnection)

	if err != nil {
		fmt.Println("db err: ", err)
		return nil, err
	}

	db.DB().SetMaxIdleConns(*conf.DBMaxIdleConns)
	db.DB().SetMaxOpenConns(*conf.DBMaxOpenConns)
	db.LogMode(true)

	return db, nil
}

func getDBConnection(conf *Config) (string, string) {
	dbType := "postgres"
	dbConnection :=
		"host=" + *conf.DBHost +
			" port=" + strconv.Itoa(*conf.DBPort) +
			" user=" + *conf.DBUser +
			" password=" + *conf.DBPassword +
			" dbname=" + *conf.DBName +
			" sslmode=" + *conf.DBSSLMode

	return dbType, dbConnection
}

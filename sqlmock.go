package sqlmock

import (
	"database/sql"
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Config config
type Config struct {
	*gorm.Config

	DriverName string
	DSN        string

	// Special configuration for mysql dialector
	SkipInitializeWithVersion bool
	DefaultStringSize         uint
	DisableDatetimePrecision  bool
	DontSupportRenameIndex    bool
	DontSupportRenameColumn   bool

	conn *sql.DB
	mock sqlmock.Sqlmock
}

// Open mock database with dsn
func Open(driverName string, dsn string) (db *gorm.DB, mock sqlmock.Sqlmock, err error) {

	config := &Config{
		DriverName: driverName,
	}

	config.conn, config.mock, err = sqlmock.NewWithDSN(dsn)
	if err != nil {
		return
	}

	return newMock(config)
}

// New mock database with config
func New(config Config) (db *gorm.DB, mock sqlmock.Sqlmock, err error) {

	config.conn, config.mock, err = sqlmock.New()
	if err != nil {
		return
	}

	return newMock(&config)
}

func newMock(config *Config) (db *gorm.DB, mock sqlmock.Sqlmock, err error) {

	var dialector gorm.Dialector

	switch config.DriverName {
	case "mysql":
		dialector = mysql.New(mysql.Config{
			DSN:                       config.DSN,
			Conn:                      config.conn,
			SkipInitializeWithVersion: config.SkipInitializeWithVersion,
			DefaultStringSize:         config.DefaultStringSize,
			DisableDatetimePrecision:  config.DisableDatetimePrecision,
			DontSupportRenameIndex:    config.DontSupportRenameIndex,
			DontSupportRenameColumn:   config.DontSupportRenameColumn,
		})

	case "postgres", "pg":
		// TODO
	case "sqlserver":
		// TODO
	case "sqlite":
		// TODO
	default:
		err = fmt.Errorf("the %s driver could not be matched", config.DriverName)
		return
	}

	db, err = gorm.Open(dialector, config.Config)

	return db, config.mock, err
}

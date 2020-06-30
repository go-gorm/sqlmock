package main

import (
	"testing"

	goSqlmock "github.com/DATA-DOG/go-sqlmock"
	"gorm.io/sqlmock"
)

func TestMigrate(t *testing.T) {

	var db, mock, err = sqlmock.New(sqlmock.Config{DriverName: "mysql"})

	if err != nil {
		t.Error("errors happened when new sqlmock: ", err.Error())
	}

	mock.ExpectExec("CREATE TABLE `posts` (.+)").WillReturnResult(goSqlmock.NewResult(1, 1))

	err = Migrate(db)
	if err != nil {
		t.Error("errors happened when migrate: ", err.Error())
	}
}

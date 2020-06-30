# sqlmock
GORM SQL Mock driver

## Install
```shell
go get gorm.io/sqlmock
```

## Something you may want to test
```go
package main

import (
	"gorm.io/gorm"
)

// Post model
type Post struct {
	gorm.Model
	Title string
	Body  string
}

// Service for post model
type Service struct {
	db *gorm.DB
}

// Create post
func (service *Service) Create(post *Post) error {

	return service.db.Create(post).Error
}
```

## Tests with sqlmock
```go
package main

import (
	"testing"
	"time"

	goSqlmock "github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"
	"gorm.io/sqlmock"
)

var service Service
var mock goSqlmock.Sqlmock

func init() {

	var (
		config = sqlmock.Config{DriverName: "mysql"}
		db     *gorm.DB
		err    error
	)

	db, mock, err = sqlmock.New(config)
	if err != nil {
		panic(err.Error())
	}

	service = Service{db: db}
}

func TestCreate(t *testing.T) {

	mock.ExpectExec("INSERT INTO `posts`").
		WithArgs(
			sqlmock.AnyTime{},   // CreatedAt
			sqlmock.AnyTime{},   // UpdatedAt
			nil,                 // DeletedAt
			"Post title",        // Title
			sqlmock.AnyString{}, // Body
		).
		WillReturnResult(goSqlmock.NewResult(1, 1))

	var post = &Post{
		Title: "Post title",
		Body:  "Post body",
	}

	err := service.Create(post)
	if err != nil {
		t.Error("errors happened when create post: ", err.Error())
	}
}
```
More [examples](https://github.com/go-gorm/sqlmock/tree/master/examples)

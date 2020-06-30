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

func TestGet(t *testing.T) {

	mock.ExpectQuery("SELECT (.+) FROM `posts` WHERE (.+) AND `posts`.`deleted_at` IS NULL ORDER BY `posts`.`id` LIMIT 1").
		WithArgs(1).
		WillReturnRows(
			goSqlmock.
				NewRows([]string{"id", "title", "body", "created_at", "updated_at", "deleted_at"}).
				AddRow(1, "Post title", "Post body", time.Now(), time.Now(), nil),
		)

	_, err := service.Get(uint(1))
	if err != nil {
		t.Error("errors happened when get post: ", err.Error())
	}
}

func TestUpdate(t *testing.T) {

	mock.ExpectExec("UPDATE `posts` SET (.+) WHERE (.+)").
		WithArgs(
			"Post body[update]",  // Body
			"Post title[update]", // Title
			sqlmock.AnyTime{},    // UpdatedAt
			1,
		).
		WillReturnResult(goSqlmock.NewResult(1, 1))

	var post = &Post{
		Model: gorm.Model{
			ID: 1,
		},
	}

	var data = map[string]interface{}{
		"title": "Post title[update]",
		"body":  "Post body[update]",
	}

	err := service.Update(post, data)
	if err != nil {
		t.Error("errors happened when update post: ", err.Error())
	}
}

func TestDestroy(t *testing.T) {

	mock.ExpectExec("UPDATE `posts` SET (.+) WHERE id = ?").
		WithArgs(sqlmock.AnyTime{}, 1).
		WillReturnResult(goSqlmock.NewResult(0, 1))

	err := service.Destroy(uint(1))
	if err != nil {
		t.Error("errors happened when delete post: ", err.Error())
	}
}

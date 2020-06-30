package sqlmock

import (
	"database/sql/driver"
	"time"
)

// Any go-sqlmock any type SQL query matching
type Any struct{}

// Match satisfies sqlmock.Argument interface
func (a Any) Match(v driver.Value) bool {
	return true
}

// AnyString go-sqlmock String type SQL query matching
type AnyString struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyString) Match(v driver.Value) bool {
	_, ok := v.(string)
	return ok
}

// AnyInt64 go-sqlmock Int64 type SQL query matching
type AnyInt64 struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyInt64) Match(v driver.Value) bool {
	_, ok := v.(int64)
	return ok
}

// AnyTime go-sqlmock Time type SQL query matching
type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

// AnyByteArray go-sqlmock []byte type SQL query matching
type AnyByteArray struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyByteArray) Match(v driver.Value) bool {
	_, ok := v.([]byte)
	return ok
}

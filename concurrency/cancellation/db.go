package cancellation

import (
	"context"
	"fmt"
	"log"
)

// -----------------------------------------------------------------------------
// Your db instance (that uses the EmulatedDriver)
// -----------------------------------------------------------------------------

// YourDB is your custom database instance that would use the EmulatedDriver.
type YourDB struct {
	driver EmulatedDriver
}

// NewYourDB creates a new instance of YourDB with the provided driver.
func NewYourDB(driver EmulatedDriver) *YourDB {
	return &YourDB{driver: driver}
}

// QueryContext is your implementation of the database query method that
// supports context cancellation.
func (db *YourDB) QueryContext(ctx context.Context, query string, args ...interface{}) (*simulatedRows, error) {

}

package arango

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/arangodb/go-driver"
	"log"
)

func ensureDatabase(ctx context.Context, c driver.Client, name string, options *driver.CreateDatabaseOptions) driver.Database {
	db, err := c.Database(ctx, name)
	if driver.IsNotFound(err) {
		db, err = c.CreateDatabase(ctx, name, options)
		if err != nil {
			if driver.IsConflict(err) {
				log.Fatalf("Failed to create database (conflict) '%s': %s %#v", name, describe(err), err)
			} else {
				log.Fatalf("Failed to create database '%s': %s %#v", name, describe(err), err)
			}
		}
	} else if err != nil {
		log.Fatalf("Failed to open database '%s': %s", name, describe(err))
	}
	return db
}

func ensureCollection(ctx context.Context, db driver.Database, name string, options *driver.CreateCollectionOptions) driver.Collection {
	c, err := db.Collection(ctx, name)
	if driver.IsNotFound(err) {
		c, err = db.CreateCollection(ctx, name, options)
		if err != nil {
			log.Fatalf("Failed to create collection '%s': %s", name, describe(err))
		}
	} else if err != nil {
		log.Fatalf("Failed to open collection '%s': %s", name, describe(err))
	}
	return c
}

// describe returns a string description of the given error.
func describe(err error) string {
	if err == nil {
		return "nil"
	}
	cause := driver.Cause(err)
	var msg string
	if re, ok := cause.(*driver.ResponseError); ok {
		msg = re.Error()
	} else {
		c, _ := json.Marshal(cause)
		msg = string(c)
	}
	if cause.Error() != err.Error() {
		return fmt.Sprintf("%v caused by %v (%v)", err, cause, msg)
	}
	return fmt.Sprintf("%v (%v)", err, msg)
}

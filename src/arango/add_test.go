package arango

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"testing"
)

type Person struct {
	Name string		`json:"name"`
	Age  int		`json:"age"`
}

func TestAdd(t *testing.T) {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://localhost:8529"},
		TLSConfig: &tls.Config{ /*...*/ },
	})
	if err != nil {
		// Handle error
	}
	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication("feng", "hdf"),
	})
	if err != nil {
		// Handle error
	}

	ctx := context.Background()

	db := ensureDatabase(ctx, client, "hdf_test", nil)

	// Create some indexes with de-duplication off
	col := ensureCollection(ctx, db, "hdf_test1_c", nil)

	docs := []map[string]interface{}{
		{ "id": 100, "name": "John", "age": 37, "active": true, "gender": "m" },
		{ "id": 101, "name": "Fred", "age": 36, "active": true, "gender": "m" },
		{ "id": 102, "name": "Jacob", "age": 35, "active": false, "gender": "m" },
		{ "id": 103, "name": "Ethan", "age": 34, "active": false, "gender": "m" },
		{ "id": 104, "name": "Michael", "age": 33, "active": true, "gender": "m" },
		{ "id": 105, "name": "Alexander", "age": 32, "active": true, "gender": "m" },
		{ "id": 106, "name": "Daniel", "age": 31, "active": true, "gender": "m" },
		{ "id": 107, "name": "Anthony", "age": 30, "active": true, "gender": "m" },
		{ "id": 108, "name": "Jim", "age": 29, "active": true, "gender": "m" },
		{ "id": 109, "name": "Diego", "age": 28, "active": true, "gender": "m" },
		{ "id": 200, "name": "Sophia", "age": 37, "active": true, "gender": "f" },
		{ "id": 201, "name": "Emma", "age": 36,  "active": true, "gender": "f" },
		{ "id": 202, "name": "Olivia", "age": 35, "active": false, "gender": "f" },
		{ "id": 203, "name": "Madison", "age": 34, "active": true, "gender": "f" },
		{ "id": 204, "name": "Chloe", "age": 33, "active": true, "gender": "f" },
		{ "id": 205, "name": "Eva", "age": 32, "active": false, "gender": "f" },
		{ "id": 206, "name": "Abigail", "age": 31, "active": true, "gender": "f" },
		{ "id": 207, "name": "Isabella", "age": 30, "active": true, "gender": "f" },
		{ "id": 208, "name": "Mary", "age": 29, "active": true, "gender": "f" },
		{ "id": 209, "name": "Mariah", "age": 28, "active": true, "gender": "f" },
	}
	meta, errs, err := col.CreateDocuments(ctx, docs)
	if driver.IsConflict(err) {
		t.Errorf("Expected Conflict error, got %s", describe(err))
	}

	fmt.Println(err)
	fmt.Println(errs)
	fmt.Println(meta)
}

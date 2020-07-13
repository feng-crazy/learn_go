package arango

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"testing"
)

func TestIndex(t *testing.T) {
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
	col := ensureCollection(ctx, db, "hdf_test_c", nil)
	if _, _, err := col.EnsureHashIndex(ctx, []string{"_key"}, &driver.EnsureHashIndexOptions{
		Unique:        true,
		Sparse:        false,
		NoDeduplicate: false,
	}); err != nil {
		t.Fatalf("Failed to create new index: %s", describe(err))
	}

	doc := struct {
		Tags []string `json:"tags"`
	}{
		Tags: []string{"e", "e"},
	}
	meta, err := col.CreateDocument(ctx, doc)
	if driver.IsConflict(err) {
		t.Errorf("Expected Conflict error, got %s", describe(err))
	}

	fmt.Println(err)
	fmt.Println(meta)
}

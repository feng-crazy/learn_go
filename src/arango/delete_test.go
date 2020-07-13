package arango

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"testing"
)

func TestDelete(t *testing.T) {
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

	query := `FOR d IN hdf_test1_c FILTER d.name=="Emma"  REMOVE d IN hdf_test1_c LET removed = OLD RETURN removed._key`
	bindVar := map[string]interface{}{}
	cursor, err := db.Query(ctx, query,bindVar)
	if err != nil{
		t.Error(err)
	}
	var doc interface{}
	_, err = cursor.ReadDocument(ctx, &doc)
	fmt.Println(err)
	fmt.Println(doc)
	if err != nil{
		t.Error(err)
	}
}
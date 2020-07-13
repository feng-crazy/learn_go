package arango

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"testing"
)

func TestQuery(t *testing.T){
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://localhost:8529"},
		TLSConfig: &tls.Config{ /*...*/ },
	})
	if err != nil {
		t.Error(err)
	}
	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication("feng", "hdf"),
	})
	if err != nil {
		t.Error(err)
	}

	ctx := context.Background()

	db := ensureDatabase(ctx, client, "hdf_test", nil)

	query := "FOR u IN users\n  UPDATE u WITH { gender: TRANSLATE(u.gender, { m: 'male', f: 'female' }) } IN users"
	bindVar := map[string]interface{}{
		"name":"%"+"h"+"%",
		"age": 18,
	}

	cursor, err := db.Query(ctx, query,bindVar)
	if err != nil{
		t.Error(err)
	}
	var doc interface{}
	_, err = cursor.ReadDocument(ctx, &doc)
	fmt.Println(err)
	fmt.Println(doc)
}

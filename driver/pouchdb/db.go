package pouchdb

import (
	"bytes"
	"context"

	"github.com/flimzy/kivik/driver/ouchdb"
	"github.com/flimzy/kivik/driver/pouchdb/bindings"
)

type db struct {
	db *bindings.DB
}

func (d *db) AllDocsContext(_ context.Context, docs interface{}, options map[string]interface{}) (offset, totalrows int, updateSeq string, err error) {
	body, err := d.db.AllDocs(options)
	if err != nil {
		return 0, 0, "", err
	}
	return ouchdb.AllDocs(bytes.NewReader(body), docs)
}

func (d *db) GetContext(_ context.Context, docID string, doc interface{}, options map[string]interface{}) error {
	return nil
}

func (d *db) CreateDocContext(_ context.Context, doc interface{}) (docID, rev string, err error) {
	return "", "", nil
}

func (d *db) PutContext(_ context.Context, docID string, doc interface{}) (rev string, err error) {
	return "", nil
}

func (d *db) DeleteContext(_ context.Context, docID, rev string) (newRev string, err error) {
	return "", nil
}

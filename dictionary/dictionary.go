package dictionary

import (
	"fmt"
	"time"

	"github.com/dgraph-io/badger/v3"
)

type Dictionary struct {
	db *badger.DB
}

type Entry struct {
	Word       string
	Definition string
	CreatedAt  time.Time
}

func (e Entry) String() string {
	created := e.CreatedAt.Format(time.Stamp)
	return fmt.Sprintf("%-10v\t%-50v%-6v", e.Word, e.Definition, created)
}

func New(dir string) (*Dictionary, error) {
	// Open the Badger database located in the /tmp/badger directory.
	// It will be created if it doesn't exist.
	db, err := badger.Open(badger.DefaultOptions("/tmp/badger"))
	if err != nil {
		return nil, err
	}
	dict := &Dictionary{
		db: db,
	}
	return dict, nil
}

func (d *Dictionary) Close() {
	d.db.Close()
}

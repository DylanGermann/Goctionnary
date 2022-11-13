package dictionary

import (
	"bytes"
	"encoding/gob"
	"time"

	"github.com/dgraph-io/badger/v3"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func (d *Dictionary) Add(word string, definition string) error {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	entry := Entry{
		Word:       cases.Title(language.Und, cases.NoLower).String(word),
		Definition: definition,
		CreatedAt:  time.Now(),
	}
	enc.Encode(entry)

	return d.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(word), buffer.Bytes())
	})
}

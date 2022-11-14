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

func (d *Dictionary) Get(word string) (Entry, error) {
	var entry Entry
	err := d.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(word))
		if err != nil {
			return err
		}
		entry, err = getEntry(item)
		return err
	})
	return entry, err
}

func getEntry(item *badger.Item) (Entry, error) {
	var entry Entry
	var buffer bytes.Buffer
	item.Value(func(val []byte) error {
		_, err := buffer.Write(val)
		return err
	})
	dec := gob.NewDecoder(&buffer)
	err := dec.Decode(&entry)
	return entry, err
}

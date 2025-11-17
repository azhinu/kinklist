package storage

import (
	"go.etcd.io/bbolt"
)

const bucketName = "kinklists"

type DB struct {
    db *bbolt.DB
}

func NewDB(path string) (*DB, error) {
    db, err := bbolt.Open(path, 0600, nil)
    if err != nil {
        return nil, err
    }

    err = db.Update(func(tx *bbolt.Tx) error {
        _, err := tx.CreateBucketIfNotExists([]byte(bucketName))
        return err
    })
    if err != nil {
        return nil, err
    }

    return &DB{db: db}, nil
}

func (d *DB) Close() error {
    return d.db.Close()
}

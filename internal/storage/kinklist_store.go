package storage

import (
	"encoding/json"
	"errors"

	"github.com/azhinu/kinklist/internal/model"
	"go.etcd.io/bbolt"
)

var ErrNotFound = errors.New("not found")

func (d *DB) GetKinkList(id string) (*model.KinkList, error) {
    var kl *model.KinkList

    err := d.db.View(func(tx *bbolt.Tx) error {
        bucket := tx.Bucket([]byte(bucketName))
        raw := bucket.Get([]byte(id))
        if raw == nil {
            return ErrNotFound
        }

        var tmp model.KinkList
        if err := json.Unmarshal(raw, &tmp); err != nil {
            return err
        }

        kl = &tmp
        return nil
    })

    return kl, err
}

func (d *DB) SaveKinkList(kl *model.KinkList) error {
    raw, err := json.Marshal(kl)
    if err != nil {
        return err
    }

    return d.db.Update(func(tx *bbolt.Tx) error {
        bucket := tx.Bucket([]byte(bucketName))
        return bucket.Put([]byte(kl.ID), raw)
    })
}

// Package badger3 provides a configurable BadgerDB v3 go-datastore for use with
// IPFS Cluster.
package badger3

import (
	"os"

	ds "github.com/ipfs/go-datastore"
	badgerds "github.com/ipfs/go-ds-badger3"
	logging "github.com/ipfs/go-log/v2"
	"github.com/pkg/errors"
)

var logger = logging.Logger("badger3")

// New returns a BadgerDB datastore configured with the given
// configuration.
func New(cfg *Config) (ds.Datastore, error) {
	folder := cfg.GetFolder()
	err := os.MkdirAll(folder, 0700)
	if err != nil {
		return nil, errors.Wrap(err, "creating badger folder")
	}
	opts := badgerds.Options{
		GcDiscardRatio: cfg.GCDiscardRatio,
		GcInterval:     cfg.GCInterval,
		GcSleep:        cfg.GCSleep,
		Options:        cfg.BadgerOptions,
	}
	return badgerds.NewDatastore(folder, &opts)
}

// Cleanup deletes the badger datastore.
func Cleanup(cfg *Config) error {
	folder := cfg.GetFolder()
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		return nil
	}
	return os.RemoveAll(cfg.GetFolder())
}

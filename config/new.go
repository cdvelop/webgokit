package config

import (
	"path/filepath"
	"time"

	bolt "go.etcd.io/bbolt"
)

var h *Handler

type Handler struct {
	storageFolder string // ej .git
	dbPath        string // ej .git/webgokit.config.db
	log           logger

	db *bolt.DB

	getter
	setter

	bucket       string // ej  config
	folderInput  string
	folderOutput string
}

type logger interface {
	Success(message string)
	Error(error)
}

func New(l logger) (*Handler, error) {

	storageFolder := ".git"
	dbPath := filepath.Join(storageFolder, "webgokit.db")

	h = &Handler{
		storageFolder: storageFolder,
		dbPath:        dbPath,
		log:           l,

		getter: getter{},
		setter: setter{},

		bucket:       "config",
		folderInput:  "folderInput",
		folderOutput: "folderOutput",
	}

	err := h.checkAndCreateStorageFolder()
	if err != nil {
		return nil, err
	}

	// db con tiempo de bloqueo en espera a 1 segundo
	h.db, err = bolt.Open(h.dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}

	// crear bucket si no existe
	err = h.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(h.bucket))
		return err
	})

	return h, err

}

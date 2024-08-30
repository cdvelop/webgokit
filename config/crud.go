package config

import (
	"fmt"

	bolt "go.etcd.io/bbolt"
)

func Create(key, value string) {
	err := h.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(h.bucket))
		if bucket == nil {
			return fmt.Errorf("bucket no encontrado al crear key: %v valor: %v", key, value)
		}
		return bucket.Put([]byte(key), []byte(value))
	})

	h.log.Error(err)
}

func Read(key string) string {
	var value []byte
	err := h.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(h.bucket))
		if bucket == nil {
			return fmt.Errorf("bucket no encontrado al leer: %v", key)
		}
		value = bucket.Get([]byte(key))
		// omitir el error solo si la app nuca se a iniciado no buk
		// if value == nil {
		// 	return fmt.Errorf("clave %v no encontrada", key)
		// }
		return nil
	})

	h.log.Error(err)

	return string(value)
}

func Update(key, newValue string) {
	err := h.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(h.bucket))
		if bucket == nil {
			return fmt.Errorf("bucket no encontrado al actualizar: %v con el valor %v", key, newValue)
		}
		return bucket.Put([]byte(key), []byte(newValue))
	})

	h.log.Error(err)
}

func Delete(key string) {
	err := h.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(h.bucket))
		if bucket == nil {
			return fmt.Errorf("error bucket no encontrado al borrar: %v", key)
		}
		return bucket.Delete([]byte(key))
	})

	h.log.Error(err)
}

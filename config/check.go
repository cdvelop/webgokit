package config

import (
	"errors"
	"os"
)

func (h *Handler) checkAndCreateStorageFolder() error {

	// Check if .git folder exists
	if _, err := os.Stat(h.storageFolder); os.IsNotExist(err) {
		// .git folder doesn't exist, create it
		err := os.Mkdir(h.storageFolder, 0755)
		if err != nil {
			// Handle error if unable to create folder
			return errors.New("Failed to create " + h.storageFolder + " folder: " + err.Error())
		}
	}

	return nil
}

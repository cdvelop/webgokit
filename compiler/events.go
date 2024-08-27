package compiler

import "path/filepath"

func (h *handler) UpdateStaticFrontendFileOnDisk(eventName string) (resetWaitingTime bool, err error) {

	h.Info("Event: " + eventName)

	if isDir(eventName) {
		// fmt.Println("Folder Event:", eventName)
	} else {

		extension := filepath.Ext(eventName)
		// fmt.Println("extension:", extension, "File Event:", eventName)

		switch extension {
		case ".css":
			// err = w.CSS.UpdateFileOnDisk(eventName)
			// if err == nil {
			// 	resetWaitingTime = true
			// }
		case ".js":
			// err = w.JS.UpdateFileOnDisk(eventName)
			// if err == nil {
			// 	resetWaitingTime = true
			// }
		case ".html":
			// err = w.HTML.UpdateFileOnDisk(eventName)

			// if err == nil {
			// 	resetWaitingTime = true
			// }

		case ".go":

			resetWaitingTime = true
		}

	}

	return
}

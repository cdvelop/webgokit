package console

import (
	"fmt"
	"os"
)

func (h *handler) Write(p []byte) (n int, err error) {

	if len(p) == 0 {
		return 0, nil
	}

	f, err := os.OpenFile(h.logFilePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		e := "Archivo " + h.logFilePath + " no existe" + err.Error()
		fmt.Println(e)
		return 0, fmt.Errorf(e)
	}
	defer f.Close()

	if len(p) != 0 {
		_, err = f.Write(p)
		if err != nil {
			e := "no se puede escribir en el archivo " + h.logFilePath + err.Error()
			fmt.Println(e)
			return 0, fmt.Errorf(e)
		}
	}

	return len(p), nil
}

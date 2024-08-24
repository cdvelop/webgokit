package gui

func (h *handler) logToConsole(message string) {
	h.console.SetText(h.console.Text + message + "\n")
	h.console.CursorRow = len(h.console.Text)
}

func (h *handler) Write(p []byte) (n int, err error) {
	h.logToConsole("\n" + string(p))
	return len(p), nil
}

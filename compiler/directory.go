package compiler

func (h *handler) InputDirectoryChange(newDirectory string) {

	h.Info("InputDirectoryChange: " + newDirectory)

}
func (h *handler) OutputDirectoryChange(newDirectory string) {

	h.Info("OutputDirectoryChange: " + newDirectory)
}

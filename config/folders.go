package config

func (s *setter) FolderInput(new string) {
	Create(h.folderInput, new)
}
func (g *getter) FolderInput() string {
	return Read(h.folderInput)
}

func (s *setter) FolderOutput(new string) {
	Create(h.folderOutput, new)
}

func (g *getter) FolderOutput() string {
	return Read(h.folderOutput)
}

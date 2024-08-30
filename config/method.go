package config

func (h *Handler) Get() *getter {
	return &getter{}
}

type getter struct{}

func (h *Handler) Set() *setter {
	return &setter{}
}

type setter struct{}

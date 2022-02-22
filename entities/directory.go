package entities

type Directory struct {
	Path string `yaml:"path"`
	Perm uint32 `yaml:"perm"`
}

func NewDirectory() *Directory {
	return &Directory{}
}

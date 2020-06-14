package manifest

type Endpoint struct {
	Path        string   `yaml:"path"`
	Description string   `yaml:"description"`
	Methods     []Method `json:"methods"`
}

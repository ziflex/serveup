package manifest

type Program struct {
	Name   string   `yaml:"name"`
	Args   []string `yaml:"args"`
	Stdin  string   `yaml:"stdin"`
	Stdout string   `yaml:"stdout"`
	Stderr string   `yaml:"stderr"`
}

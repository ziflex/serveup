package manifest

type Method struct {
	Name        string  `yaml:"name"`
	Description string  `yaml:"description"`
	Program     Program `yaml:"program"`
}

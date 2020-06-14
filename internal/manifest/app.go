package manifest

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type Application struct {
	Name      string `yaml:"name"`
	Version   string `yaml:"version"`
	Endpoints []Endpoint
}

func Unmarshal(content []byte) (Application, error) {
	app := Application{}

	if err := yaml.Unmarshal(content, &app); err != nil {
		return app, errors.Wrap(err, "failed to parse app manifest")
	}

	return app, nil
}

func Marshal(app Application) ([]byte, error) {
	return yaml.Marshal(app)
}

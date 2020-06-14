package runtime

import (
	"github.com/Masterminds/sprig/v3"
	"html/template"
)

func NewTemplate(name string, text string) (*template.Template, error) {
	return template.New(name).Funcs(sprig.FuncMap()).Parse(text)
}

package runtime

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"os/exec"

	"github.com/pkg/errors"

	"github.com/ziflex/serveup/internal/manifest"
)

type Program struct {
	name   string
	args   []*template.Template
	stdin  *template.Template
	stdout *template.Template
	stderr *template.Template
}

func NewProgram(config manifest.Program) (Program, error) {
	program := Program{
		args: make([]*template.Template, len(config.Args)),
	}

	for i, a := range config.Args {
		arg, err := NewTemplate(fmt.Sprintf("%s.args.%d", config.Name, i), a)

		if err != nil {
			return Program{}, errors.Wrapf(err, "parse args: %d", i)
		}

		program.args[i] = arg
	}

	if len(config.Stdin) > 0 {
		stdin, err := NewTemplate(fmt.Sprintf("%s.stdin", config.Name), config.Stdin)

		if err != nil {
			return Program{}, errors.Wrap(err, "parse stdin")
		}

		program.stdin = stdin
	}

	if len(config.Stdout) > 0 {
		stdout, err := NewTemplate(fmt.Sprintf("%s.stdout", config.Name), config.Stdout)

		if err != nil {
			return Program{}, errors.Wrap(err, "parse stdout")
		}

		program.stdout = stdout
	}

	if len(config.Stdout) > 0 {
		stderr, err := NewTemplate(fmt.Sprintf("%s.stderr", config.Name), config.Stderr)

		if err != nil {
			return Program{}, errors.Wrap(err, "parse stderr")
		}

		program.stderr = stderr
	}

	return program, nil
}

func (p *Program) Exec(ctx context.Context) ([]byte, error) {
	args := make([]string, len(p.args))

	runtimeCtx := NewContext(ctx.Value("env"), ctx.Value("req"))

	for i, a := range p.args {
		var b bytes.Buffer

		if err := a.Execute(&b, runtimeCtx); err != nil {
			return nil, errors.Wrapf(err, "execute argument template: %d", i)
		}

		args[i] = b.String()
	}

	cmd := exec.CommandContext(ctx, p.name, args...)

	if p.stdin != nil {
		stdin := &bytes.Buffer{}

		if err := p.stdin.Execute(stdin, runtimeCtx); err != nil {
			return nil, errors.Wrap(err, "execute stdin template")
		}

		cmd.Stdin = stdin
	}

	out, err := cmd.CombinedOutput()

	if err != nil {
		if p.stderr == nil {
			return nil, err
		}

		stderr := &bytes.Buffer{}
		runtimeCtx.Exec.Error = err.Error()

		if err := p.stderr.Execute(stderr, runtimeCtx); err != nil {
			return nil, errors.Wrap(err, "execute stderr template")
		}

		runtimeCtx.Exec.Error = stderr.String()

		return nil, errors.New(runtimeCtx.Exec.Error)
	}

	if p.stdout == nil {
		return out, nil
	}

	stdout := &bytes.Buffer{}
	runtimeCtx.Exec.Result = string(out)

	if err := p.stdout.Execute(stdout, runtimeCtx); err != nil {
		return nil, errors.Wrap(err, "execute stdout template")
	}

	return stdout.Bytes(), nil
}

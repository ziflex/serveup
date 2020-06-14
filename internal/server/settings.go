package server

import (
	"github.com/ziflex/serveup/internal/manifest"
)

type Settings struct {
	Version  string
	Port     uint64
	Manifest manifest.Application
}

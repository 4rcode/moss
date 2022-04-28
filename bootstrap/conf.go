package bootstrap

import (
	cli "github.com/4rcode/moss/cliconf"
	"github.com/4rcode/moss/conf"
	env "github.com/4rcode/moss/envconf"
	file "github.com/4rcode/moss/fileconf"
	json "github.com/4rcode/moss/jsonconf"
)

// Configuration TODO
type Configuration struct {
	Cli  cli.Configurer
	Env  env.Configurer
	File file.ConfigurerFactory
	Json json.ConfigurerFactory
}

// Configure TODO
func (c Configuration) Configure(value interface{}) error {
	if c.File.Factory == nil {
		c.File.Factory = c.Json
	}

	if c.Env.StringFactory == nil {
		c.Env.StringFactory = c.File
	}

	if c.Env.ReaderFactory == nil {
		c.Env.ReaderFactory = c.Json
	}

	if c.Cli.Factories.File == nil {
		c.Cli.Factories.File = c.File
	}

	if c.Cli.Factories.Inline == nil {
		c.Cli.Factories.Inline = c.Json
	}

	return conf.Configurers{
		c.Env, c.Cli,
	}.Configure(value)
}

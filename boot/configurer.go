package boot

import (
	"github.com/4rcode/moss/conf"
	"github.com/4rcode/moss/conf/cli"
	"github.com/4rcode/moss/conf/env"
	"github.com/4rcode/moss/conf/file"
	"github.com/4rcode/moss/conf/json"
)

// Configurer TODO
type Configurer struct {
	Configurers Configurers
	Factories   ConfigurerFactories
}

// Configurers TODO
type Configurers struct {
	Cli cli.Configurer
	Env env.Configurer
}

// ConfigurerFactories TODO
type ConfigurerFactories struct {
	File file.ConfigurerFactory
	Json json.ConfigurerFactory
}

// Configure TODO
func (c Configurer) Configure(value interface{}) error {
	if c.Factories.File.ConfigurerFactory == nil {
		c.Factories.File.ConfigurerFactory = c.Factories.Json
	}

	if c.Configurers.Env.Factories.File == nil {
		c.Configurers.Env.Factories.File = c.Factories.File
	}

	if c.Configurers.Env.Factories.Inline == nil {
		c.Configurers.Env.Factories.Inline = c.Factories.Json
	}

	if c.Configurers.Cli.Factories.File == nil {
		c.Configurers.Cli.Factories.File = c.Factories.File
	}

	if c.Configurers.Cli.Factories.Inline == nil {
		c.Configurers.Cli.Factories.Inline = c.Factories.Json
	}

	return conf.Configurers{
		c.Configurers.Env, c.Configurers.Cli,
	}.Configure(value)
}

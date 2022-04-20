package boot

import (
	"github.com/4rcode/moss/conf"
	"github.com/4rcode/moss/conf/cli"
	"github.com/4rcode/moss/conf/env"
	"github.com/4rcode/moss/conf/file"
	"github.com/4rcode/moss/conf/json"
)

// Configurer TODO
type Configurer struct{}

func (Configurer) Configure(value interface{}) error {
	return conf.Configurers{
		env.Configurer{
			PathConfigurerFactory:   file.ConfigurerFactory{},
			ReaderConfigurerFactory: json.ConfigurerFactory{},
		},
		cli.Configurer{
			PathConfigurerFactory:   file.ConfigurerFactory{},
			ReaderConfigurerFactory: json.ConfigurerFactory{},
		},
	}.Configure(value)
}

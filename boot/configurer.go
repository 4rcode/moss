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
	var jsonConfigurerFactory json.ConfigurerFactory

	fileConfigurerFactory := file.ConfigurerFactory{
		ConfigurerFactory: jsonConfigurerFactory,
	}

	return conf.Configurers{
		env.Configurer{
			FileConfigurerFactory:   fileConfigurerFactory,
			InlineConfigurerFactory: jsonConfigurerFactory,
		},
		cli.Configurer{
			FileConfigurerFactory:   fileConfigurerFactory,
			InlineConfigurerFactory: jsonConfigurerFactory,
		},
	}.Configure(value)
}

package boot

import (
	"github.com/4rcode/moss/conf"
	"github.com/4rcode/moss/conf/cli"
	"github.com/4rcode/moss/conf/env"
)

// Configurer TODO
type Configurer struct{}

func (Configurer) Configure(value interface{}) error {
	return conf.Configurers{
		env.Configurer{},
		cli.Configurer{},
	}.Configure(value)
}

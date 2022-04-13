package boot

import (
	"github.com/4rcode/moss/conf"
	"github.com/4rcode/moss/conf/cli"
	"github.com/4rcode/moss/conf/env"
)

// Configuration TODO
type Configuration struct{}

func (Configuration) Parse(value any) error {
	return conf.Parsers{
		env.Parser{},
		cli.Parser{},
	}.Parse(value)
}

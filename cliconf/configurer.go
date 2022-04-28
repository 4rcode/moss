package cli

import (
	"flag"
	"strings"

	"github.com/4rcode/moss/conf"
)

// Configurer TODO
type Configurer struct {
	Factories ConfigurerFactories
	Flags     Flags
	FlagSet   *flag.FlagSet
	Labels    UsageLabels
}

// ConfigurerFactories TODO
type ConfigurerFactories struct {
	File   conf.StringFactory
	Inline conf.ReaderFactory
}

// Flags TODO
type Flags struct {
	File, Inline string
}

// UsageLabels TODO
type UsageLabels Flags

// Configure TODO
func (c Configurer) Configure(value interface{}) error {
	if c.Flags.File == "" {
		c.Flags.File = "c"
	}

	if c.Flags.Inline == "" {
		c.Flags.Inline = "C"
	}

	if c.FlagSet == nil {
		c.FlagSet = flag.CommandLine
	}

	if c.Factories.File != nil {
		c.FlagSet.Var(_flag{
			c.buildFileConfigurer, value,
		}, c.Flags.File, c.Labels.File)
	}

	if c.Factories.Inline != nil {
		c.FlagSet.Var(_flag{
			c.buildInlineConfigurer, value,
		}, c.Flags.Inline, c.Labels.Inline)
	}

	return nil
}

func (c Configurer) buildFileConfigurer(text string) conf.Configurer {
	return c.Factories.File.NewConfigurer(text)
}

func (c Configurer) buildInlineConfigurer(text string) conf.Configurer {
	return c.Factories.Inline.NewConfigurer(
		strings.NewReader(text),
	)
}

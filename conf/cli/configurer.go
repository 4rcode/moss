package cli

import (
	"flag"
	"io"
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
	File interface {
		NewConfigurer(...string) conf.Configurer
	}

	Inline interface {
		NewConfigurer(io.Reader) conf.Configurer
	}
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

	var file, inline string

	c.FlagSet.StringVar(&file, c.Flags.File, file, c.Labels.File)
	c.FlagSet.StringVar(&inline, c.Flags.Inline, inline, c.Labels.Inline)

	if file != "" && c.Factories.File != nil {
		err := c.Factories.File.
			NewConfigurer(file).
			Configure(value)

		if err != nil {
			return err
		}
	}

	if inline != "" && c.Factories.Inline != nil {
		err := c.Factories.Inline.
			NewConfigurer(
				strings.NewReader(inline)).
			Configure(value)

		if err != nil {
			return err
		}
	}

	return nil
}

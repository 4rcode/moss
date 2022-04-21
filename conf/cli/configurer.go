package cli

import (
	"flag"
	"io"
	"strings"

	"github.com/4rcode/moss/conf"
)

// Configurer TODO
type Configurer struct {
	// Flags TODO
	Flags Flags

	// FlagSet TODO
	FlagSet *flag.FlagSet

	// Labels TODO
	Labels Labels

	// FileConfigurerFactory TODO
	FileConfigurerFactory interface {
		NewConfigurer(...string) conf.Configurer
	}

	// InlineConfigurerFactory TODO
	InlineConfigurerFactory interface {
		NewConfigurer(io.Reader) conf.Configurer
	}
}

type Flags struct {
	// File TODO
	File,

	// Inline TODO
	Inline string
}

type Labels struct {
	// File TODO
	File,

	// Inline TODO
	Inline string
}

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

	if file != "" {
		err := c.
			FileConfigurerFactory.
			NewConfigurer(file).
			Configure(value)

		if err != nil {
			return err
		}
	}

	if inline != "" {
		err := c.
			InlineConfigurerFactory.
			NewConfigurer(
				strings.NewReader(inline)).
			Configure(value)

		if err != nil {
			return err
		}
	}

	return nil
}

package cli

import (
	"flag"
	"io"
	"strings"

	"github.com/4rcode/moss/conf"
)

// Configurer TODO
type Configurer struct {
	// FileFlag TODO
	FileFlag,

	// FileFlagUsage TODO
	FileFlagUsage,

	// InlineFlag TODO
	InlineFlag,

	// InlineFlagUsage TODO
	InlineFlagUsage string

	// FlagSet TODO
	FlagSet *flag.FlagSet

	// PathConfigurerFactory TODO
	PathConfigurerFactory interface {
		NewConfigurer(...string) conf.Configurer
	}

	// ReaderConfigurerFactory TODO
	ReaderConfigurerFactory interface {
		NewConfigurer(io.Reader) conf.Configurer
	}
}

func (c Configurer) Configure(value interface{}) error {
	if c.FileFlag == "" {
		c.FileFlag = "c"
	}

	if c.InlineFlag == "" {
		c.InlineFlag = "C"
	}

	if c.FlagSet == nil {
		c.FlagSet = flag.CommandLine
	}

	var file, inline string

	c.FlagSet.StringVar(&file, c.FileFlag, file, c.FileFlagUsage)
	c.FlagSet.StringVar(&inline, c.InlineFlag, inline, c.InlineFlagUsage)

	if file != "" {
		err := c.
			PathConfigurerFactory.
			NewConfigurer(file).
			Configure(value)

		if err != nil {
			return err
		}
	}

	if inline != "" {
		err := c.
			ReaderConfigurerFactory.
			NewConfigurer(
				strings.NewReader(inline)).
			Configure(value)

		if err != nil {
			return err
		}
	}

	return nil
}

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

	// FileConfigurerFactory TODO
	FileConfigurerFactory interface {
		NewConfigurer(...string) conf.Configurer
	}

	// InlineConfigurerFactory TODO
	InlineConfigurerFactory interface {
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

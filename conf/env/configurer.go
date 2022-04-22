package env

import (
	"io"
	"os"
	"strings"

	"github.com/4rcode/moss/conf"
)

// Configurer TODO
type Configurer struct {
	Factories ConfigurerFactories
	Variables Variables
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

// Variables TODO
type Variables struct {
	File, Inline string
}

// Configure TODO
func (c Configurer) Configure(value interface{}) error {
	if c.Variables.File == "" {
		c.Variables.File = "MOSS_FILE"
	}

	if c.Variables.Inline == "" {
		c.Variables.Inline = "MOSS_CONF"
	}

	file := os.Getenv(c.Variables.File)
	inline := os.Getenv(c.Variables.Inline)

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

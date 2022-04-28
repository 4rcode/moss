package env

import (
	"os"
	"strings"

	"github.com/4rcode/moss/conf"
)

// Configurer TODO
type Configurer struct {
	File, Inline string

	ReaderFactory conf.ReaderFactory
	StringFactory conf.StringFactory
}

// Configure TODO
func (c Configurer) Configure(value interface{}) error {
	if c.File == "" {
		c.File = "MOSS_CONFIG_FILES"
	}

	if c.Inline == "" {
		c.Inline = "MOSS_CONFIG"
	}

	file := os.Getenv(c.File)
	inline := os.Getenv(c.Inline)

	if file != "" && c.StringFactory != nil {
		err := c.StringFactory.
			NewConfigurer(file).
			Configure(value)

		if err != nil {
			return err
		}
	}

	if inline != "" && c.ReaderFactory != nil {
		err := c.ReaderFactory.
			NewConfigurer(
				strings.NewReader(inline)).
			Configure(value)

		if err != nil {
			return err
		}
	}

	return nil
}

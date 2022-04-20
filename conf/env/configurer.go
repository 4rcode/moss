package env

import (
	"fmt"
	"io"
	"os"

	"github.com/4rcode/moss/conf"
)

// Configurer TODO
type Configurer struct {
	// FileVar TODO
	FileVar,

	// InlineVar TODO
	InlineVar string

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
	if c.FileVar == "" {
		c.FileVar = "MOSS_FILE"
	}

	if c.InlineVar == "" {
		c.InlineVar = "MOSS_CONF"
	}

	fileVar := os.Getenv(c.FileVar)
	inlineVar := os.Getenv(c.InlineVar)

	fmt.Println(fileVar, inlineVar)

	return nil
}

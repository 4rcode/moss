package env

import (
	"fmt"
	"os"
)

// Configurer TODO
type Configurer struct {
	FileVar, InlineVar string
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

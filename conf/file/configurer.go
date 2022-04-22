package file

import (
	"io"
	"os"
	"strings"

	"github.com/4rcode/moss/conf"
)

// Configurer TODO
type Configurer struct {
	ConfigurerFactory interface {
		NewConfigurer(io.Reader) conf.Configurer
	}

	Paths     []string
	Separator string
}

// Configure TODO
func (c Configurer) Configure(value interface{}) error {
	if value == nil || c.ConfigurerFactory == nil {
		return nil
	}

	if c.Separator == "" {
		c.Separator = string(os.PathListSeparator)
	}

	for _, paths := range c.Paths {
		for _, path := range strings.Split(paths, c.Separator) {
			file, err := os.Open(path)

			if os.IsNotExist(err) {
				continue
			}

			if err != nil {
				return err
			}

			defer file.Close()

			err = c.
				ConfigurerFactory.
				NewConfigurer(file).
				Configure(value)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

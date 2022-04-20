package file

import (
	"io"
	"os"
	"strings"

	"github.com/4rcode/moss/conf"
)

// Configurer TODO
type Configurer struct {
	// Path TODO
	Path []string

	// ReaderParserFactory TODO
	ReaderParserFactory func(io.Reader) conf.Configurer

	// Separator TODO
	Separator string
}

func (c Configurer) Parse(value interface{}) error {
	if c.ReaderParserFactory == nil {
		return nil
	}

	if c.Separator == "" {
		c.Separator = string(os.PathListSeparator)
	}

	for _, paths := range c.Path {
		for _, path := range strings.Split(paths, c.Separator) {
			file, err := os.Open(path)

			if err != nil {
				return err
			}

			defer file.Close()

			if err = c.
				ReaderParserFactory(file).
				Configure(value); err != nil {
				return err
			}
		}
	}

	return nil
}

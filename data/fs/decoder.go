package fsdata

import (
	"os"
	"strings"

	"github.com/4rcode/moss/data"
)

// Decoder TODO
type Decoder struct {
	Builder data.ReaderDecoderBuilder

	Paths     []string
	Separator string
}

// Decode TODO
func (c Decoder) Decode(value interface{}) error {
	if value == nil || c.Builder == nil {
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
				Builder.
				Build(file).
				Decode(value)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

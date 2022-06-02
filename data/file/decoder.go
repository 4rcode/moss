package file

import (
	"errors"
	"io/fs"
	"os"
	"strings"

	"github.com/4rcode/moss/data"
)

// Decoder TODO
type Decoder struct {
	Builder data.DecoderBuilder

	Path, Separator string
}

// Decode TODO
func (d Decoder) Decode(value interface{}) error {
	if value == nil || d.Builder == nil {
		return nil
	}

	if d.Separator == "" {
		d.Separator = string(os.PathListSeparator)
	}

	for _, path := range strings.Split(d.Path, d.Separator) {
		file, err := os.Open(path)

		if errors.Is(err, fs.ErrNotExist) {
			continue
		}

		if err != nil {
			return err
		}

		defer file.Close()

		err = d.
			Builder.
			Build(file).
			Decode(value)

		if err != nil {
			return err
		}
	}

	return nil
}

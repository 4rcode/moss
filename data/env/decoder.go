package env

import (
	"os"
	"strings"

	"github.com/4rcode/moss/data"
)

// Decoder TODO
type Decoder struct {
	File, Inline string

	ReaderBuilder, StringBuilder data.DecoderBuilder
}

// Decode TODO
func (d Decoder) Decode(value interface{}) error {
	if d.File == "" {
		d.File = "MOSS_CONFIG_FILES"
	}

	if d.Inline == "" {
		d.Inline = "MOSS_CONFIG"
	}

	file := os.Getenv(d.File)
	inline := os.Getenv(d.Inline)

	if file != "" && d.StringBuilder != nil {
		err := d.
			StringBuilder.
			Build(
				strings.NewReader(file)).
			Decode(value)

		if err != nil {
			return err
		}
	}

	if inline != "" && d.ReaderBuilder != nil {
		err := d.ReaderBuilder.Build(
			strings.NewReader(inline),
		).Decode(value)

		if err != nil {
			return err
		}
	}

	return nil
}

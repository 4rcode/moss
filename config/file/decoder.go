package conf

import (
	"io"
	"os"
	"strings"

	"github.com/4rcode/moss/config"
	"github.com/4rcode/moss/config/json"
	"github.com/4rcode/moss/internal"
)

// FileParser TODO
type FileParser struct {
	// Path TODO
	Path []string

	// ReaderParserFactory TODO
	ReaderParserFactory func(io.Reader) config.Decoder

	// Separator TODO
	Separator string
}

func (p FileParser) Parse(target interface{}) error {
	if p.ReaderParserFactory == nil {
		p.ReaderParserFactory = p.newJsonParser
	}

	internal.Set(&p.Separator, "", string(os.PathListSeparator))

	for _, paths := range p.Path {
		for _, path := range strings.Split(paths, p.Separator) {
			file, err := os.Open(path)

			if err != nil {
				return err
			}

			defer file.Close()

			if err = p.
				ReaderParserFactory(file).
				Decode(target); err != nil {
				return err
			}
		}
	}

	return nil
}

func (p *FileParser) newJsonParser(reader io.Reader) config.Decoder {
	return json.Decoder{Reader: reader}
}

package json

import (
	"encoding/json"
	"io"

	"github.com/4rcode/moss/conf"
)

// Configurer TODO
type Configurer struct {
	Decoder *json.Decoder
}

// Configure TODO
func (c Configurer) Configure(value interface{}) error {
	if value == nil || c.Decoder == nil {
		return nil
	}

	for {
		err := c.Decoder.Decode(value)

		if err == nil {
			continue
		}

		if err == io.EOF {
			return nil
		}

		return err
	}
}

// ConfigurerFactory TODO
type ConfigurerFactory func(io.Reader) *json.Decoder

// NewConfigurer TODO
func (f ConfigurerFactory) NewConfigurer(reader io.Reader) conf.Configurer {
	if reader == nil {
		return Configurer{}
	}

	if f == nil {
		f = json.NewDecoder
	}

	return Configurer{
		f(reader),
	}
}

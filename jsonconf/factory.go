package jsonconf

import (
	"encoding/json"
	"io"
)

// Factory TODO
type Factory func(io.Reader) *json.Decoder

// NewConfigurer TODO
func (f Factory) NewConfigurer(reader io.Reader) Configurer {
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

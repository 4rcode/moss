package jsonprop

import (
	"encoding/json"
	"io"
)

// Factory TODO
type Factory func(io.Reader) *json.Decoder

// NewConfigurer TODO
func (f Factory) NewConfigurer(reader io.Reader) Decoder {
	if reader == nil {
		return Decoder{}
	}

	if f == nil {
		f = json.NewDecoder
	}

	return Decoder{
		f(reader),
	}
}

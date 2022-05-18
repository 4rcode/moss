package jsondata

import (
	"encoding/json"
	"io"

	"github.com/4rcode/moss/data"
)

// DecoderBuilder TODO
type DecoderBuilder func(io.Reader) *json.Decoder

// Build TODO
func (b DecoderBuilder) Build(reader io.Reader) data.Decoder {
	if reader == nil {
		return Decoder{}
	}

	if b == nil {
		b = json.NewDecoder
	}

	return Decoder{
		b(reader),
	}
}

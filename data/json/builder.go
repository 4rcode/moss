package json

import (
	"encoding/json"
	"io"

	"github.com/4rcode/moss/data"
	"github.com/4rcode/moss/data/seq"
)

// DecoderBuilder TODO
type DecoderBuilder func(io.Reader) *json.Decoder

// Build TODO
func (b DecoderBuilder) Build(reader io.Reader) data.Decoder {
	if reader == nil {
		return seq.Decoder{}
	}

	if b == nil {
		b = json.NewDecoder
	}

	return seq.
		DecoderBuilder[*json.Decoder](b).
		Build(reader)
}

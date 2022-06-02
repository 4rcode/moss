package xml

import (
	"encoding/xml"
	"io"

	"github.com/4rcode/moss/data"
	"github.com/4rcode/moss/data/seq"
)

// DecoderBuilder TODO
type DecoderBuilder func(io.Reader) *xml.Decoder

// Build TODO
func (b DecoderBuilder) Build(reader io.Reader) data.Decoder {
	if reader == nil {
		return seq.Decoder{}
	}

	if b == nil {
		b = xml.NewDecoder
	}

	return seq.
		DecoderBuilder[*xml.Decoder](b).
		Build(reader)
}

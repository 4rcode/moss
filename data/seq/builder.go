package seq

import (
	"io"

	"github.com/4rcode/moss/data"
)

// DecoderBuilder TODO
type DecoderBuilder[D data.Decoder] func(io.Reader) D

// Build TODO
func (b DecoderBuilder[D]) Build(reader io.Reader) Decoder {
	if b == nil || reader == nil {
		return Decoder{}
	}

	return Decoder{
		Decoder: b(reader),
	}
}

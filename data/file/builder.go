package file

import (
	"io"

	"github.com/4rcode/moss/data"
)

// DecoderBuilder TODO
type DecoderBuilder Decoder

// Build TODO
func (b DecoderBuilder) Build(reader io.Reader) data.Decoder {
	if reader != nil {
		if bytes, err := io.ReadAll(reader); err == nil {
			b.Path = string(bytes)
		}
	}

	return Decoder(b)
}

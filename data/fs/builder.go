package fsdata

import "github.com/4rcode/moss/data"

// DecoderBuilder TODO
type DecoderBuilder Decoder

// Build TODO
func (b DecoderBuilder) Build(paths ...string) data.Decoder {
	b.Paths = append(b.Paths, paths...)

	return Decoder(b)
}

package data

import (
	"io"
)

// StringDecoderBuilder TODO
type StringDecoderBuilder interface {
	// Build TODO
	Build(...string) Decoder
}

// ReaderDecoderBuilder TODO
type ReaderDecoderBuilder interface {
	// Build TODO
	Build(io.Reader) Decoder
}

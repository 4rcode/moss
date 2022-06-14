package seq

import (
	"io"

	"github.com/4rcode/moss/data"
)

// Decoder TODO
type Decoder struct {
	Decoder data.Decoder
}

// Decode TODO
func (d Decoder) Decode(value interface{}) error {
	if value == nil || d.Decoder == nil {
		return nil
	}

	for {
		err := d.Decoder.Decode(value)

		if err == nil {
			continue
		}

		if err == io.EOF {
			return nil
		}

		return err
	}
}

// DecoderBuilder TODO
type DecoderBuilder[D data.Decoder] func(io.Reader) D

// Build TODO
func (b DecoderBuilder[D]) Build(reader io.Reader) data.Decoder {
	if b == nil || reader == nil {
		return Decoder{}
	}

	return Decoder{b(reader)}
}

package json

import (
	"encoding/json"
	"io"
)

// ConfigurerFactory TODO
type ConfigurerFactory struct {
	DecoderFactory interface {
		NewDecoder(io.Reader) *json.Decoder
	}
}

func (c ConfigurerFactory) NewConfigurer(reader io.Reader) *Configurer {
	if reader == nil {
		return nil
	}

	var decoder *json.Decoder

	if c.DecoderFactory != nil {
		decoder = c.DecoderFactory.NewDecoder(reader)
	}

	if decoder == nil {
		decoder = c.NewDecoder(reader)
	}

	return &Configurer{decoder}
}

func (c ConfigurerFactory) NewDecoder(reader io.Reader) *json.Decoder {
	if reader == nil {
		return &json.Decoder{}
	}

	return json.NewDecoder(reader)
}

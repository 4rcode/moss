package json

import (
	"encoding/json"
	"io"

	"github.com/4rcode/moss/conf"
)

// ConfigurerFactory TODO
type ConfigurerFactory struct {
	DecoderFactory interface {
		NewDecoder(io.Reader) *json.Decoder
	}
}

func (c ConfigurerFactory) NewConfigurer(reader io.Reader) conf.Configurer {
	if reader == nil {
		return &Configurer{}
	}

	var decoder *json.Decoder

	if c.DecoderFactory == nil {
		decoder = json.NewDecoder(reader)
	} else {
		decoder = c.DecoderFactory.NewDecoder(reader)
	}

	return &Configurer{decoder}
}

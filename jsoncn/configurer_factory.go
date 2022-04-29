package jsoncn

import (
	"encoding/json"
	"io"
)

// ConfigurerFactory TODO
type ConfigurerFactory func(io.Reader) *json.Decoder

// NewConfigurer TODO
func (f ConfigurerFactory) NewConfigurer(reader io.Reader) Configurer {
	if reader == nil {
		return Configurer{}
	}

	if f == nil {
		f = json.NewDecoder
	}

	return Configurer{
		f(reader),
	}
}

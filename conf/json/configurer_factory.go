package json

import (
	"encoding/json"
	"io"

	"github.com/4rcode/moss/conf"
)

// ConfigurerFactory TODO
type ConfigurerFactory func(io.Reader) *json.Decoder

// NewConfigurer TODO
func (f ConfigurerFactory) NewConfigurer(reader io.Reader) conf.Configurer {
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

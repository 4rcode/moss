package json

import (
	"encoding/json"
	"io"

	"github.com/4rcode/moss/conf"
)

// ConfigurerFactory TODO
type ConfigurerFactory struct{}

// NewConfigurer TODO
func (c ConfigurerFactory) NewConfigurer(reader io.Reader) conf.Configurer {
	var decoder *json.Decoder

	if reader != nil {
		decoder = json.NewDecoder(reader)
	}

	return &Configurer{decoder}
}

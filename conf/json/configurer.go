package json

import (
	"encoding/json"
	"io"
)

// Configurer TODO
type Configurer struct {
	Decoder *json.Decoder
}

func (c Configurer) Configure(value interface{}) error {
	if value == nil || c.Decoder == nil {
		return nil
	}

	for {
		err := c.Decoder.Decode(value)

		if err == nil {
			continue
		}

		if err == io.EOF {
			return nil
		}

		return err
	}
}

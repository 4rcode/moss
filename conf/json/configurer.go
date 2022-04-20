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
	decoder := c.Decoder

	if value == nil || decoder == nil {
		return nil
	}

	for {
		err := decoder.Decode(value)

		if err == nil {
			continue
		}

		if err == io.EOF {
			return nil
		}

		return err
	}
}

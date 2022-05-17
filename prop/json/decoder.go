package jsonprop

import (
	"encoding/json"
	"io"
)

// Decoder TODO
type Decoder struct {
	Decoder *json.Decoder
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

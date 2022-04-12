package json

import (
	"encoding/json"
	"io"
)

// Decoder
type Decoder struct {
	Reader io.Reader
}

func (d Decoder) Decode(value any) error {
	decoder := json.NewDecoder(d.Reader)

	for {
		if err := decoder.Decode(value); err != nil {
			if err == io.EOF {
				return nil
			}

			return err
		}
	}
}

package json

import (
	"encoding/json"
	"io"
)

// Parser
type Parser struct {
	Reader io.Reader
}

func (d Parser) Parse(value any) error {
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

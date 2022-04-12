package config

// Decoders TODO
type Decoders []Decoder

func (d Decoders) Decode(value any) error {
	for _, decoder := range d {
		if decoder == nil {
			continue
		}

		if err := decoder.Decode(value); err != nil {
			return err
		}
	}

	return nil
}

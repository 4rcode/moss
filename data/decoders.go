package data

// Decoders TODO
type Decoders []Decoder

// Decode TODO
func (d Decoders) Decode(value interface{}) error {
	for _, decoder := range d {
		if decoder == nil {
			continue
		}

		err := decoder.Decode(value)

		if err != nil {
			return err
		}
	}

	return nil
}

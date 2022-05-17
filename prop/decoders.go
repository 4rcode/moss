package prop

type Decoders []Decoder

func (d Decoders) Configure(value interface{}) error {
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

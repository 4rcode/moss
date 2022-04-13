package conf

// Parsers TODO
type Parsers []Parser

func (d Parsers) Parse(value any) error {
	for _, decoder := range d {
		if decoder == nil {
			continue
		}

		if err := decoder.Parse(value); err != nil {
			return err
		}
	}

	return nil
}

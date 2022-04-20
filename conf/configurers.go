package conf

// Parsers TODO
type Configurers []Configurer

func (c Configurers) Configure(value interface{}) error {
	for _, configurer := range c {
		if configurer == nil {
			continue
		}

		if err := configurer.Configure(value); err != nil {
			return err
		}
	}

	return nil
}

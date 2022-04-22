package conf

// Configurers TODO
type Configurers []Configurer

func (c Configurers) Configure(value interface{}) error {
	for _, configurer := range c {
		if configurer == nil {
			continue
		}

		err := configurer.Configure(value)

		if err != nil {
			return err
		}
	}

	return nil
}

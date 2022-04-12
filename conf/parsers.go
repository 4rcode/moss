package conf

// Parse TODO
func Parse(target interface{}) error {
	return Parsers{
		EnvParser{},
		CliParser{},
	}.Parse(target)
}

// Parser TODO
type Parser interface {
	// Parse TODO
	Parse(interface{}) error
}

// Parsers TODO
type Parsers []Parser

func (p Parsers) Parse(target interface{}) error {
	return nil
}

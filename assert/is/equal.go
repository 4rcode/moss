package is

import "fmt"

type Equal struct {
	To     interface{}
	Format string
}

func (e Equal) Match(value interface{}) error {
	if e.To == value {
		return nil
	}

	return fmt.Errorf("%v %v %v", value, "is not", e.To)
}

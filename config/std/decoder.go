package std

import (
	"github.com/4rcode/moss/config"
	"github.com/4rcode/moss/config/cli"
	"github.com/4rcode/moss/config/env"
)

//Decoder TODO
type Decoder struct{}

func (Decoder) Parse(target any) error {
	return config.Decoders{
		env.Decoder{},
		cli.Decoder{},
	}.Decode(target)
}

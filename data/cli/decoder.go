package clidata

import (
	"flag"
	"strings"

	"github.com/4rcode/moss/data"
)

// Decoder TODO
type Decoder struct {
	Factories ConfigurerFactories
	Flags     Flags
	FlagSet   *flag.FlagSet
	Labels    UsageLabels
}

// ConfigurerFactories TODO
type ConfigurerFactories struct {
	File   data.StringDecoderBuilder
	Inline data.ReaderDecoderBuilder
}

// Flags TODO
type Flags struct {
	File, Inline string
}

// UsageLabels TODO
type UsageLabels Flags

// Decode TODO
func (c Decoder) Decode(value interface{}) error {
	if c.Flags.File == "" {
		c.Flags.File = "c"
	}

	if c.Flags.Inline == "" {
		c.Flags.Inline = "C"
	}

	if c.FlagSet == nil {
		c.FlagSet = flag.CommandLine
	}

	if c.Factories.File != nil {
		c.FlagSet.Var(_flag{
			c.buildFileConfigurer, value,
		}, c.Flags.File, c.Labels.File)
	}

	if c.Factories.Inline != nil {
		c.FlagSet.Var(_flag{
			c.buildInlineConfigurer, value,
		}, c.Flags.Inline, c.Labels.Inline)
	}

	return nil
}

func (c Decoder) buildFileConfigurer(text string) data.Decoder {
	return c.Factories.File.Build(text)
}

func (c Decoder) buildInlineConfigurer(text string) data.Decoder {
	return c.Factories.Inline.Build(
		strings.NewReader(text),
	)
}

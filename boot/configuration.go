package boot

import (
	"encoding/json"

	"github.com/4rcode/moss/data"
	"github.com/4rcode/moss/data/cli"
	"github.com/4rcode/moss/data/env"
	"github.com/4rcode/moss/data/file"
	"github.com/4rcode/moss/data/seq"
)

// Configuration TODO
type Configuration struct {
	Cli    cli.Decoder
	Env    env.Decoder
	File   file.DecoderBuilder
	Inline data.DecoderBuilder
}

// Configure TODO
func (c Configuration) Configure(value interface{}) error {
	if c.Inline == nil {
		c.Inline = seq.DecoderBuilder[*json.Decoder](json.NewDecoder)
	}

	if c.File.Builder == nil {
		c.File.Builder = c.Inline
	}

	if c.Env.StringBuilder == nil {
		c.Env.StringBuilder = c.File
	}

	if c.Env.ReaderBuilder == nil {
		c.Env.ReaderBuilder = c.File
	}

	if c.Cli.Factories.File == nil {
		c.Cli.Factories.File = c.File
	}

	if c.Cli.Factories.Inline == nil {
		c.Cli.Factories.Inline = c.File
	}

	return data.Decoders{
		c.Env, c.Cli,
	}.Decode(value)
}

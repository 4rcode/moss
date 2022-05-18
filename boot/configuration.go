package boot

import (
	"github.com/4rcode/moss/data"
	cli "github.com/4rcode/moss/data/cli"
	env "github.com/4rcode/moss/data/env"
	fs "github.com/4rcode/moss/data/fs"
	json "github.com/4rcode/moss/data/json"
)

// Configuration TODO
type Configuration struct {
	Cli  cli.Decoder
	Env  env.Decoder
	File fs.DecoderBuilder
	Json json.DecoderBuilder
}

// Configure TODO
func (c Configuration) Configure(value interface{}) error {
	if c.File.Builder == nil {
		c.File.Builder = c.Json
	}

	if c.Env.StringBuilder == nil {
		c.Env.StringBuilder = c.File
	}

	if c.Env.ReaderBuilder == nil {
		c.Env.ReaderBuilder = c.Json
	}

	if c.Cli.Factories.File == nil {
		c.Cli.Factories.File = c.File
	}

	if c.Cli.Factories.Inline == nil {
		c.Cli.Factories.Inline = c.Json
	}

	return data.Decoders{
		c.Env, c.Cli,
	}.Decode(value)
}

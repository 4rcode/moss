package cli

import (
	"flag"

	"github.com/4rcode/moss/internal"
)

type Parser struct {
	FileEnvVar, FileFlag,
	InlineEnvVar, InlineFlag string

	FlagSet *flag.FlagSet
}

func (p Parser) Parse(target interface{}) {
	internal.Set(&p.FileFlag, "", "c")
	internal.Set(&p.InlineFlag, "", "C")
	internal.Set(&p.FileEnvVar, "", "c")
	internal.Set(&p.InlineEnvVar, "", "C")
	internal.Set(&p.FlagSet, nil, flag.CommandLine)
}

func Parse(target interface{}) {
	Parser{}.Parse(target)
}

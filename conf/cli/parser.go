package cli

import (
	"flag"
	"fmt"

	"github.com/4rcode/moss/internal"
)

type Parser struct {
	FileFlag, FileFlagUsage,
	InlineFlag, InlineFlagUsage string

	FlagSet *flag.FlagSet
}

func (p Parser) Parse(target interface{}) error {
	internal.Set(&p.FileFlag, "", "c")
	internal.Set(&p.InlineFlag, "", "C")
	internal.Set(&p.FlagSet, nil, flag.CommandLine)

	var file, inline string

	p.FlagSet.StringVar(&file, p.FileFlag, file, p.FileFlagUsage)
	p.FlagSet.StringVar(&inline, p.InlineFlag, inline, p.InlineFlagUsage)

	fmt.Println("")

	return nil
}

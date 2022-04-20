package cli

import (
	"flag"
	"fmt"
)

// Configurer TODO
type Configurer struct {
	FileFlag, FileFlagUsage,
	InlineFlag, InlineFlagUsage string

	FlagSet *flag.FlagSet
}

func (c Configurer) Configure(target interface{}) error {
	if c.FileFlag == "" {
		c.FileFlag = "c"
	}

	if c.InlineFlag == "" {
		c.InlineFlag = "C"
	}

	if c.FlagSet == nil {
		c.FlagSet = flag.CommandLine
	}

	var file, inline string

	c.FlagSet.StringVar(&file, c.FileFlag, file, c.FileFlagUsage)
	c.FlagSet.StringVar(&inline, c.InlineFlag, inline, c.InlineFlagUsage)

	fmt.Println("")

	return nil
}

package cli

import (
	"flag"
	"fmt"

	"github.com/4rcode/moss/internal"
)

type Decoder struct {
	FileFlag, FileFlagUsage,
	InlineFlag, InlineFlagUsage string

	FlagSet *flag.FlagSet
}

func (d Decoder) Decode(target interface{}) error {
	internal.Set(&d.FileFlag, "", "c")
	internal.Set(&d.InlineFlag, "", "C")
	internal.Set(&d.FlagSet, nil, flag.CommandLine)

	var file, inline string

	d.FlagSet.StringVar(&file, d.FileFlag, file, d.FileFlagUsage)
	d.FlagSet.StringVar(&inline, d.InlineFlag, inline, d.InlineFlagUsage)

	fmt.Println("")

	return nil
}

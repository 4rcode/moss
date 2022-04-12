package conf

import (
	"fmt"
	"os"

	"github.com/4rcode/moss/internal"
)

type EnvParser struct {
	FileVar, InlineVar string
}

func (p EnvParser) Parse(target interface{}) error {
	internal.Set(&p.FileVar, "", "c")
	internal.Set(&p.InlineVar, "", "C")

	fileVar := os.Getenv(p.FileVar)
	inlineVar := os.Getenv(p.InlineVar)

	fmt.Println(fileVar, inlineVar)

	return nil
}

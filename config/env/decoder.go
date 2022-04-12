package env

import (
	"fmt"
	"os"

	"github.com/4rcode/moss/internal"
)

type Decoder struct {
	FileVar, InlineVar string
}

func (d Decoder) Decode(value any) error {
	internal.Set(&d.FileVar, "", "c")
	internal.Set(&d.InlineVar, "", "C")

	fileVar := os.Getenv(d.FileVar)
	inlineVar := os.Getenv(d.InlineVar)

	fmt.Println(fileVar, inlineVar)

	return nil
}

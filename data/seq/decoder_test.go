package seq_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/4rcode/moss/data"
	"github.com/4rcode/moss/data/seq"
)

func TestConfigurer(t *testing.T) {
	type _data struct {
		Foo string
	}

	var myData _data

	data.Decoder(seq.Decoder{
		json.NewDecoder(strings.NewReader(""))},
	).Decode(&myData)

	t.Error(myData)
}

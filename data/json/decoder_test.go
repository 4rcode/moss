package jsondata_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/4rcode/moss/data"
	jsondata "github.com/4rcode/moss/data/json"
)

func TestConfigurer(t *testing.T) {
	type _data struct {
		Foo string
	}

	var myData _data

	data.Decoder(jsondata.Decoder{
		json.NewDecoder(strings.NewReader(""))},
	).Decode(&myData)

	t.Error(myData)
}

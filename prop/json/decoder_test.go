package jsonprop_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/4rcode/moss/property"
	jsonprop "github.com/4rcode/moss/property/json"
)

func TestConfigurer(t *testing.T) {
	type _data struct {
		Foo    string
		Foobar struct {
			foo, Bar string
		}
	}

	var data _data

	property.Decoder(jsonprop.Decoder{
		json.NewDecoder(strings.NewReader(""))},
	).Decode(&data)

	t.Error(data)
}

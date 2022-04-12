package json_test

import (
	"strings"
	"testing"

	"github.com/4rcode/moss/config"
	"github.com/4rcode/moss/config/json"
)

func TestAbc(t *testing.T) {
	type _data struct {
		Foo    string
		bar    int
		Foobar struct {
			foo, Bar string
		}
	}

	var data _data

	config.Decoder(json.Decoder{
		strings.NewReader(""),
	}).Decode(&data)

	t.Error(data)
}

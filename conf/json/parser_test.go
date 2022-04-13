package json_test

import (
	"strings"
	"testing"

	"github.com/4rcode/moss/conf"
	"github.com/4rcode/moss/conf/json"
)

func TestParser(t *testing.T) {
	type _data struct {
		Foo    string
		bar    int
		Foobar struct {
			foo, Bar string
		}
	}

	var data _data

	conf.Parser(json.Parser{
		strings.NewReader(""),
	}).Parse(&data)

	t.Error(data)
}

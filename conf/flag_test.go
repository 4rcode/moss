package conf_test

import (
	"encoding/json"
	"testing"
)

func TestXxx(t *testing.T) {
	type bla struct {
		Abc int
	}

	var a bla

	json.Unmarshal([]byte(`{"Abc": 123}`), &a)

	t.Error(a)
}

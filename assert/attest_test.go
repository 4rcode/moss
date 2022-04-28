package assert_test

import (
	"testing"

	"github.com/4rcode/moss/assert"
)

func TestError(t *testing.T) {
	is := assert.Error(t).Attest

	is(1 == 2, "no it's not")
}

package is_test

import (
	"testing"

	"github.com/4rcode/moss/assert/is"
)

func TestXxx(t *testing.T) {
	t.Error(123, is.Equal{To: 123}.Match("123"))
}

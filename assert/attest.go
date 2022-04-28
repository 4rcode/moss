package assert

import (
	"testing"
)

func Error(tb testing.TB) Attester {
	return Attester{tb.Error}
}

func Fatal(tb testing.TB) Attester {
	return Attester{tb.Fatal}
}

func Log(tb testing.TB) Attester {
	return Attester{tb.Log}
}

func Skip(tb testing.TB) Attester {
	return Attester{tb.Skip}
}

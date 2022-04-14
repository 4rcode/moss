package attesting

type Printer func(string, ...any)

// Attester TODO
type Attester []Printer

func (a Attester) Attest(truth bool, args ...any) bool {
	if truth {
		return true
	}

	for _, print := range a {
		if print != nil {
			print("FAIL")
		}
	}

	return false
}

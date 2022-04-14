package attesting

// Attester TODO
type Attester []func(...any)

// Attest TODO
func (a Attester) Attest(truth bool, args ...any) bool {
	if truth {
		return true
	}

	for _, log := range a {
		if log != nil {
			log(args...)
		}
	}

	return false
}

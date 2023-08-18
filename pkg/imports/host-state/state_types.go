package host_state

type Metadata map[string]string

type GetOption struct {
}

type SetOption struct {
}

//nolint:nosnakecase
func stateConsistencyToString(c uint32) string { //nolint:nosnakecase
	switch c {
	case 1:
		return "eventual"
	case 2:
		return "strong"
	}

	return ""
}

//nolint:nosnakecase
func stateConcurrencyToString(c int) string {
	switch c {
	case 1:
		return "first-write"
	case 2:
		return "last-write"
	}

	return ""
}

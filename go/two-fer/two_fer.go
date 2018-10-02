package twofer

// ShareWith returns "One for X, one for me.", where X is the given name, or "you", if name is empty.
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	}
	return "One for " + name + ", one for me."
}

package helper

// FindFirstDifference finds the first difference between two string arrays.
// It returns the first element of array "a" that is not present in array "b",
// or an empty string if all elements of "a" are present in "b".
func FindFirstDifference(a, b []string) string {
	bMap := make(map[string]bool, len(b))
	for _, v := range b {
		bMap[v] = true
	}

	for _, v := range a {
		if !bMap[v] {
			return v
		}
	}

	return ""
}

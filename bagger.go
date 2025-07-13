package jsonbagger

// ExtractJSON returns the first JSON object found in the input string.
func ExtractJSON(input string) (string, error) {
	begin, end, err := extractJSONIndexes(input)
	if err != nil {
		return "", err
	}

	return input[begin:end], nil
}

// extractJSONIndexes returns the indexes of the first JSON object found in the input string.
func extractJSONIndexes(input string) (begin, end int, err error) {
	var jsonFound bool
	var count uint8
	for i, character := range input {
		if character == '{' {
			if count == 255 {
				err = ErrNestingOverflow
				return
			}
			count++
			if !jsonFound {
				begin = i
			}
			jsonFound = true
		}
		if character == '}' && count > 0 {
			count--
		}
		if count == 0 && jsonFound {
			end = i + 1
			return
		}
	}

	err = ErrNotFound
	return
}

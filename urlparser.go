package urlparser

import "strings"

// RemoveURLQueries
//
// If no queryKeys is inserted, it will automatically assume to remove every instance of URL query.
func RemoveURLQueries(inputURL string, queryKeys ...string) string {
	// Find URL query separator index.
	separatorIdx := strings.Index(inputURL, "?")
	if separatorIdx == -1 { // no URL queries
		return inputURL
	}

	// Keep URL fragment if any.
	urlFragment := ""
	fragmentHashIdx := strings.Index(inputURL, "#")
	if fragmentHashIdx != -1 {
		urlFragment = inputURL[fragmentHashIdx:]
	}

	// No specified query keys.
	if len(queryKeys) == 0 {
		return inputURL[:separatorIdx] + urlFragment
	}

	// Remove certain query keys.
	rawQuery := ""
	if fragmentHashIdx == -1 {
		rawQuery = inputURL[separatorIdx+1:]
	} else {
		rawQuery = inputURL[separatorIdx+1 : fragmentHashIdx]
	}
	pairs := strings.Split(rawQuery, "&")
	var kept []string
	for _, pair := range pairs {
		if hasQueryKey(pair, queryKeys) {
			continue
		}
		kept = append(kept, pair)
	}

	// No URL query is found to be kept.
	if len(kept) == 0 {
		return inputURL[:separatorIdx] + urlFragment
	}

	return inputURL[:separatorIdx] + "?" + strings.Join(kept, "&") + urlFragment
}

func hasQueryKey(pair string, queryKeys []string) bool {
	for _, key := range queryKeys {
		if strings.HasPrefix(pair, key+"=") {
			return true
		}
	}
	return false
}

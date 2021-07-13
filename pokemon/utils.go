package pokemon

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func isNotFoundError(notFoundErrors []string) plugin.ErrorPredicate {
	return func(err error) bool {
		errMsg := err.Error()
		for _, msg := range notFoundErrors {
			if strings.Contains(errMsg, msg) {
				return true
			}
		}
		return false
	}
}

func extractUrlOffset(fullUrl string) (int, error) {
	// Get the next offset number from the URL, e.g., https://pokeapi.co/api/v2/pokemon/?offset=20&limit=20
	u, err := url.Parse(fullUrl)
	if err != nil {
		return 0, err
	}

	m, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return 0, err
	}

	// Offset query param only has 1 value
	urlOffset, err := strconv.Atoi(m["offset"][0])
	if err != nil {
		return 0, err
	}

	return urlOffset, nil
}

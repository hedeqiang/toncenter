package services

import (
	"fmt"
	"net/url"
)

// buildQueryString builds query string from parameters
func buildQueryString(endpoint string, params map[string]interface{}) string {
	if len(params) == 0 {
		return endpoint
	}

	values := url.Values{}
	for key, value := range params {
		if value == nil {
			continue
		}

		switch v := value.(type) {
		case []string:
			for _, item := range v {
				values.Add(key, item)
			}
		default:
			values.Add(key, fmt.Sprintf("%v", value))
		}
	}

	query := values.Encode()
	if query != "" {
		return fmt.Sprintf("%s?%s", endpoint, query)
	}

	return endpoint
}

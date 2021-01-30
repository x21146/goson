package goson

import "strings"

// parse tag and skip other json options
func parseTag(tag string) (string, options) {
	if idx := strings.Index(tag, ","); idx != -1 {
		return tag[:idx], parseOptions(tag[idx+1:])
	}
	return tag, nil
}

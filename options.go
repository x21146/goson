package gson

import "strings"

type option string
type options []option

const (
	omitempty option = "omitempty"
)

func parseOptions(tag string) options {
	tags := strings.Split(tag, ",")
	opts := make(options, len(tags))
	for i := range opts {
		opts[i] = option(strings.TrimSpace(tags[i]))
	}
	return opts
}

func (o options) Contains(opt option) bool {
	for _, op := range o {
		if op == opt {
			return true
		}
	}
	return false
}
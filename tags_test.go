package goson

import "testing"

func TestParseTags(t *testing.T) {
	tag := "value,omitempty,test"
	tg, opt := parseTag(tag)

	if tg != "value" {
		t.FailNow()
		return
	}

	if opt[0] != omitempty && opt[1] != "test" {
		t.FailNow()
		return
	}
}

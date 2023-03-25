package hitokoto

import (
	"net/url"
)

type Request interface {
	FormatToValues() url.Values
}

type Response interface {
	Parse([]byte) error
}

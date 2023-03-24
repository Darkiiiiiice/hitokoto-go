package op

import (
	"net/url"

	"github.com/mariomang/hitokoto/constants"
)

type Response[T any] struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []T    `json:"data"`
	Ts      int64  `json:"ts"`
}

type Command interface {
	Values() url.Values
	Parse([]byte) error
	API() *constants.API
}

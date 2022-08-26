package request

import (
	"context"
	"net/http"
)

type Request interface {
	Do(context.Context, *http.Request) ([]byte, int, error)
}

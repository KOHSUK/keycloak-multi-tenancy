package server

import "net/http"

const (
	_ = iota
	GET
	POST
	PUT
	DELETE
)

type Route interface {
	http.Handler
}

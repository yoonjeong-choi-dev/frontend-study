package subcmd

import "errors"

var ErrNoServerSpecified = errors.New("have to specify the remote server")

// ErrInvalidHTTPMethod : Exercise 2.2
var ErrInvalidHTTPMethod = errors.New("invalid HTTP Method")

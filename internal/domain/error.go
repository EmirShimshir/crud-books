package domain

import "errors"

var ErrBookNotFound = errors.New("book not found")
var ErrDeleteFailed = errors.New("no deleted rows")

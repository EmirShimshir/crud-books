package domain

import "errors"

var ErrBookNotFound = errors.New("book not found")
var ErrDeleteFailed = errors.New("no deleted rows")
var ErrInvalidID = errors.New("id can't be less 1")

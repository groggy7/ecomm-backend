package domain

import "errors"

var (
	ErrProductNotFound error = errors.New("product not found")
	ErrOrderNotFound   error = errors.New("order not found")
)

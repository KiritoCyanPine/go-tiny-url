package persistant

import "errors"

var (
	ErrKeyCollision = errors.New("key collision")
	ErrKeyNotFound  = errors.New("key not found")
)

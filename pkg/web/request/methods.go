package request

type Method int

const (
	GET = iota
	POST
	PATCH
	DELETE
	UPDATE
	HEAD
)

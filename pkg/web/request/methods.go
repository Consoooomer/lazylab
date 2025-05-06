package request

type RequestMethod int

const (
	GET = iota
	POST
	PATCH
	DELETE
	UPDATE
	HEAD
)

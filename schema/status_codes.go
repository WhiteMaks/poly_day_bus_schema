package schema

type Status int64

const (
	OK             Status = 2000
	BAD_REQUEST    Status = 4000
	NOT_AUTHORIZED Status = 4010
	NOT_FOUND      Status = 4040
	INTERNAL_ERROR Status = 5000
)

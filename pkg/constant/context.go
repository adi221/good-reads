package constant

type key int

const (
	// ContextRequestID is the key used to store request ID into the request context
	ContextRequestID key = iota
)

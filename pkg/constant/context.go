package constant

type key int

const (
	// ContextRequestID is the key used to store request ID into the request context
	ContextRequestID key = iota
	// ContextCookieSetter is the key used to allow to set cookie inside service
	ContextCookieSetter
	// ContextAuthorizationTokenSetter is the key used to set authorization token in response header
	ContextAuthorizationTokenSetter
	// ContextUserID is the key used to store current user ID into the request context
	ContextUserID
	// ContextUser is the key used to store current user into the request context
	ContextUser
)

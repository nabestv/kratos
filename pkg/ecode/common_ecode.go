package ecode

// All common ecode
var (
	OK = add(0) // ok

	NotModified        = add(-304) // Wood has changed
	TemporaryRedirect  = add(-307) // Crash jump
	RequestErr         = add(-400) // Request error
	Unauthorized       = add(-401) // Unauthorized
	AccessDenied       = add(-403) // Insufficient access
	NothingFound       = add(-404) // 404
	MethodNotAllowed   = add(-405) // This method is not supported
	Conflict           = add(-409) // conflict
	Canceled           = add(-498) // Client cancel request
	ServerErr          = add(-500) // Server Error
	ServiceUnavailable = add(-503) // Overload protection, service is temporarily unavailable
	Deadline           = add(-504) // Service call timed out
	LimitExceed        = add(-509) // exceed the limit
)

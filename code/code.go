package code

type Code string

const (
	OK                   Code = "OK"
	BadRequest           Code = "BAD_REQUEST"
	Unauthorized         Code = "UNAUTHORIZED"
	PermissionDenied     Code = "PERMISSION_DENIED"
	NotFound             Code = "NOT_FOUND"
	MethodNotAllowed     Code = "METHOD_NOT_ALLOWED"
	Conflict             Code = "CONFLICT"
	UnsupportedMediaType Code = "UNSUPPORTED_MEDIA_TYPE"
	UnprocessableEntity  Code = "UNPROCESSABLE_ENTITY"
	FailedPrecondition   Code = "FAILED_PRECONDITION"
	InternalServerError  Code = "INTERNALSERVERERROR"
	Unknown              Code = "UNKNOWN"
)

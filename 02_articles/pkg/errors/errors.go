package errors

var (
	// 4xx client error
	IncorrectRequest     = new("incorrect request", 400)
	IncorrectRequestBody = new("incorrect request body", 400)

	RecordingError  = new("recording error", 400)
	NoCategoryFound = new("no categories found", 404)

	IncorrectRequestParams = new("incorrect request parametrs", 400)
	Unauthorized           = new("unauthorized", 401)
	Forbidden              = new("forbidden", 403)
	NotFound               = new("not found", 404)
	NotAllowed             = new("method not allowed", 405)
	CategoryIsAlready      = new("this category already exists", 409)

	// 5xx server error
	InternalServerError = new("internal server error", 500)
	BadGateway          = new("server received invalid response from upstream", 502)
	UnknownError        = new("unknown error", 520)
)

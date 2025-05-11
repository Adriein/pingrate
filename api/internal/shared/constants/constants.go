package constants

//Env var keys

const (
	DatabaseUser       = "DATABASE_USER"
	DatabasePassword   = "DATABASE_PASSWORD"
	DatabaseName       = "DATABASE_NAME"
	ServerPort         = "SERVER_PORT"
	Env                = "ENV"
	Production         = "PRO"
	JwtSecret          = "JWT_SECRET"
	GoogleClientId     = "GOOGLE_CLIENT_ID"
	GoogleClientSecret = "GOOGLE_CLIENT_SECRET"
)

// Criteria

const (
	Equal              = "="
	GreaterThanOrEqual = ">="
	LessThanOrEqual    = "<="
	LeftJoin           = "LEFT"
)

// Errors

const (
	ServerGenericError = "SERVER_ERROR"
	ValidationError    = "VALIDATION_ERROR"
)

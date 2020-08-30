package output

import (
	"database/sql"
	"errors"
	"net/http"
)

var (
	//ErrURLParams is the error code for bad data in the URL parameters
	ErrURLParams = errors.New("Invalid url parameters")

	ErrUserNotLogged = errors.New("User not logged in")

	//ErrInsertingData is the error code for error when inserting data on the database
	ErrInsertingData = errors.New("Error inserting data")

	//ErrInternalServer is the error code for internal server error, this is the error returned by default if none of the errors match the one used
	ErrInternalServer = errors.New("Yep, server fucked :(")

	//ErrNoContentFound is the error code for a request that has been performed correctly but has not found data on the database
	ErrNoContentFound = errors.New("No content found")

	//ErrInvalidToken is the error code for an invalid json web token
	ErrInvalidToken = errors.New("Token invalid")

	//ErrMalformedToken is the error code for a json web token with a wrong structure
	ErrMalformedToken = errors.New("Token malformed")

	//ErrMissingToken is the error code for a missing json web token on the request headers
	ErrMissingToken = errors.New("Missing auth token")

	//ErrWrongBasicAuth is the error code for a wrong authentication using basicAuth
	ErrWrongBasicAuth = errors.New("Invalid username or password")

	//ErrBodyParams is the error code for wrong parameters in the request body
	ErrBodyParams = errors.New("Invalid request payload")

	//ErrFetchResponse is the error for a response that fails using the implemented Fetch function
	ErrFetchResponse = errors.New("The fetch call did not return success")

	//ErrTooManyRequests is the error for when a user has performed too many requests and is rate limited
	ErrTooManyRequests = errors.New("Request limit exceeded, try again in a few seconds")

	//ErrTooMuchData is the error for when a user has performed too many requests and is rate limited
	ErrTooMuchData = errors.New("Too much data sent, the maximum is 500 elements per request")

	//SuccessResource is the code for when a resource is created successfully
	SuccessResource = string("Resource created correctly")
)

/*
ErrorCodeMessage returns an http error and a message for a particular error
This function is called in respondwitherror to determine the user error message
*/
func ErrorCodeMessage(_err error) (int, string) {
	switch _err {
	case sql.ErrNoRows:
		return http.StatusNoContent, "No data found"

	case ErrURLParams:
		return http.StatusBadRequest, _err.Error()

	case ErrBodyParams:
		return http.StatusBadRequest, _err.Error()

	case ErrMissingToken:
		return http.StatusUnauthorized, _err.Error()

	case ErrWrongBasicAuth:
		return http.StatusUnauthorized, _err.Error()

	case ErrMalformedToken:
		return http.StatusForbidden, _err.Error()

	case ErrInsertingData:
		return http.StatusInternalServerError, _err.Error()

	case ErrInternalServer:
		return http.StatusInternalServerError, _err.Error()

	case ErrFetchResponse:
		return http.StatusNoContent, _err.Error()

	case ErrTooManyRequests:
		return http.StatusTooManyRequests, _err.Error()

	case ErrTooMuchData:
		return http.StatusBadRequest, _err.Error()

	case ErrUserNotLogged:
		return http.StatusUnauthorized, _err.Error()

	default:
		return http.StatusInternalServerError, "Yep, server fucked :("

	}
}

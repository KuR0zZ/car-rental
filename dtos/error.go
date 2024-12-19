// @Failure      400      {object}  echo.HTTPError  "Invalid request body or data"
// @Failure      404      {object}  echo.HTTPError  "Car not available"
// @Failure      422      {object}  echo.HTTPError  "Insufficient balance"
// @Failure      500      {object}  echo.HTTPError  "Internal server error"

package dtos

type ErrorBadRequest struct {
	Message string `json:"message" example:"bad request data"`
}

type ErrorUnauthorized struct {
	Message string `json:"message" example:"invalid credential"`
}

type ErrorNotFound struct {
	Message string `json:"message" example:"data not found"`
}

type ErrorConflict struct {
	Message string `json:"message" example:"email is already registered"`
}

type ErrorUnprocessableEntity struct {
	Message string `json:"message" example:"insufficient balance"`
}

type ErrorInternalServerError struct {
	Message string `json:"message" example:"internal server error"`
}

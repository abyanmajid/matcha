package internal

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

// WriteJSON writes the given data as JSON to the provided http.ResponseWriter.
// It sets the response status code and optional headers.
//
// Parameters:
//   - w: The http.ResponseWriter to write the JSON response to.
//   - data: The data to be marshaled into JSON and written to the response.
//   - status: The HTTP status code to set for the response.
//   - headers: Optional additional headers to set on the response.
//
// Returns:
//   - error: An error if JSON marshaling or writing to the response fails.
func WriteJSON(w http.ResponseWriter, data any, status int, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

// WriteErrorJSON writes an error message in JSON format to the provided http.ResponseWriter.
// It logs the error message using slog.Debug and sets the HTTP status code to the provided status code
// or defaults to http.StatusBadRequest if no status code is provided.
//
// Parameters:
//   - w: The http.ResponseWriter to write the JSON response to.
//   - err: The error to be written in the JSON response.
//   - status: Optional variadic parameter to specify the HTTP status code.
//
// Returns:
//   - An error if there is an issue writing the JSON response.
func WriteErrorJSON(w http.ResponseWriter, err error, status ...int) error {
	slog.Debug(err.Error())

	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	response := map[string]string{
		"error": err.Error(),
	}

	return WriteJSON(w, response, statusCode)
}

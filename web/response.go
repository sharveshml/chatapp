package web

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

type httpStatus interface {
	HttpStatus() int
}

func Respond(ctx context.Context, w http.ResponseWriter, dataModel Encoder) error {

	if err := ctx.Err(); err != nil {
		if errors.Is(err, context.Canceled) {
			return errors.New("Clinet disconnected, do not send any response ")
		}
	}

	var statusCode = http.StatusOK

	switch v := dataModel.(type) {
	case httpStatus:
		statusCode = v.HttpStatus()
	case error:
		statusCode = http.StatusInternalServerError
	default:
		if dataModel == nil {
			statusCode = http.StatusNoContent
		}
	}

	if statusCode == http.StatusNoContent {
		w.WriteHeader(statusCode)
		return nil
	}

	data, contentType, err := dataModel.Encode()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return fmt.Errorf("response %w: ", err)
	}

	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)

	if _, err := w.Write(data); err != nil {
		return fmt.Errorf("Error writing the response %w: ", err)
	}

	return nil
}

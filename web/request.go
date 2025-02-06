package web

import (
	"fmt"
	"io"
	"net/http"
)

type Decoder interface {
	Decode(data []byte) error
}

type validator interface {
	Validate() error
}

func Decode(r http.Request, d Decoder) error {
	data, err := io.ReadAll(r.Body)

	if err != nil {
		return fmt.Errorf("Unable to read the request payload: %w", err)
	}

	if err := d.Decode(data); err != nil {
		return fmt.Errorf("Unable to decode the request: %w", err)
	}

	if v, err := d.(validator); err {
		if err := v.Validate(); err != nil {
			return err
		}
	}
	return nil
}

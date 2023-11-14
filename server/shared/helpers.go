package shared

import (
	"encoding/json"
	"io"
	"log"
)

func RequestBodyParser[T any](body io.ReadCloser, input T) error {
	// No body passed in
	if body == nil {
		return nil
	}

	b, err := io.ReadAll(body)
	if err != nil {
		log.Println(err)
		return err
	}

	err = json.Unmarshal(b, &input)
	if err != nil {
		log.Println(err)
	}
	return err
}

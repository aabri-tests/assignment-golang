package decoder

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"

	"github.com/aabri-assignments/assignment-golang/pkg/logger"
)

type Decoder interface {
	Decode(r io.Reader, v interface{}) error
}

type jsonDecoder struct {
	logger *logger.Logger
}

func New(logger *logger.Logger) Decoder {
	return &jsonDecoder{logger: logger}
}

func (d *jsonDecoder) Decode(r io.Reader, v interface{}) error {
	if r == nil {
		return fmt.Errorf("nil reader passed to Decode")
	}

	if v == nil {
		return fmt.Errorf("nil value passed to Decode")
	}

	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	err := dec.Decode(v)
	if err != nil {
		var syntaxErr *json.SyntaxError
		if errors.As(err, &syntaxErr) {
			return fmt.Errorf("invalid JSON input")
		}

		d.logger.Printf("error decoding object: %s. Details: %v", err, map[string]interface{}{
			"type":  reflect.TypeOf(v),
			"input": r,
		})
	}

	return err
}

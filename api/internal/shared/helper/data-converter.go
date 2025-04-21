package helper

import (
	"encoding/json"
	"github.com/rotisserie/eris"
	"io"
	"net/http"
)

func Encode[T any](w http.ResponseWriter, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(v); err != nil {
		return eris.Errorf("encode json: %v", err)
	}

	return nil
}

func Decode[T any](r io.Reader) (T, error) {
	var v T
	if err := json.NewDecoder(r).Decode(&v); err != nil {
		return v, eris.Errorf("decode json: %v", err)
	}
	return v, nil
}

package types

import "net/http"

type PingrateHttpHandler func(w http.ResponseWriter, r *http.Request) error

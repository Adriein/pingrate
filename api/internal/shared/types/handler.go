package types

import "net/http"

type Ctx struct {
	Res  http.ResponseWriter
	Req  *http.Request
	Data map[string]interface{}
}

type PingrateHttpHandler func(ctx *Ctx) error

package pingrate

import (
	"fmt"
	"github.com/adriein/pingrate/internal/shared/constants"
	"github.com/adriein/pingrate/internal/shared/container"
	"github.com/adriein/pingrate/internal/shared/helper"
	"github.com/adriein/pingrate/internal/shared/middleware"
	"github.com/adriein/pingrate/internal/shared/types"
	"github.com/rotisserie/eris"
	"log"
	"log/slog"
	"net/http"
)

type HttpServer struct {
	address   string
	router    *http.ServeMux
	container container.Container
}

func (s *HttpServer) Start() {
	v1 := http.NewServeMux()
	v1.Handle("/api/v1/", http.StripPrefix("/api/v1", s.router))

	server := http.Server{
		Addr:    s.address,
		Handler: middleware.RequestTracing(v1),
	}

	slog.Info("Starting the PingrateApiServer at " + s.address)

	err := server.ListenAndServe()

	if err != nil {
		pingrateErr := eris.Wrap(err, "Error starting HTTP server")

		log.Fatal(eris.ToString(pingrateErr, true))
	}
}

func (s *HttpServer) Route(url string, handler types.PingrateHttpHandler, middlewares ...types.Middleware) {
	s.router.Handle(url, s.httpHandler(handler, middlewares...))
}

func (s *HttpServer) httpHandler(handler types.PingrateHttpHandler, middlewares ...types.Middleware) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := make(map[string]interface{})
		data[container.ContainerInstanceKey] = s.container

		ctx := &types.Ctx{
			Res:  w,
			Req:  r,
			Data: data,
		}

		if len(middlewares) == 0 {
			if err := handler(ctx); err != nil {
				response := types.ServerResponse{
					Ok:    false,
					Error: constants.ServerGenericError,
				}

				encodeErr := helper.Encode[types.ServerResponse](w, http.StatusInternalServerError, response)

				if encodeErr != nil {
					log.Fatal(eris.ToString(encodeErr, true))
				}

				slog.Error(fmt.Sprintf("%s TraceId=%s", eris.ToString(err, true), r.Header.Get("traceId")))
			}
		}

		for i := 0; i < len(middlewares); i++ {
			if err := middlewares[i](handler)(ctx); err != nil {
				response := types.ServerResponse{
					Ok:    false,
					Error: constants.ServerGenericError,
				}

				encodeErr := helper.Encode[types.ServerResponse](w, http.StatusInternalServerError, response)

				if encodeErr != nil {
					log.Fatal(eris.ToString(encodeErr, true))
				}

				slog.Error(fmt.Sprintf("%s TraceId=%s", eris.ToString(err, true), r.Header.Get("traceId")))
			}
		}
	}
}

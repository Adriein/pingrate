package server

import (
	"fmt"
	"github.com/adriein/pingrate/internal/shared/constants"
	"github.com/adriein/pingrate/internal/shared/helper"
	"github.com/adriein/pingrate/internal/shared/middleware"
	"github.com/adriein/pingrate/internal/shared/types"
	"github.com/rotisserie/eris"
	"log"
	"log/slog"
	"net/http"
)

type PingrateApiServer struct {
	address      string
	router       *http.ServeMux
	dependencies map[string]interface{}
}

func New(address string, container map[string]interface{}) (*PingrateApiServer, error) {
	router := http.NewServeMux()

	return &PingrateApiServer{
		address:      address,
		router:       router,
		dependencies: container,
	}, nil
}

func (s *PingrateApiServer) Start() {
	v1 := http.NewServeMux()
	v1.Handle("/api/v1/", http.StripPrefix("/api/v1", s.router))

	/*MuxMiddleWareChain := middleware.NewMiddlewareChain(
		middleware.NewRequestTracingMiddleware,
	)

	server := http.Server{
		Addr:    s.address,
		Handler: MuxMiddleWareChain.ApplyOn(v1),
	}*/

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

func (s *PingrateApiServer) Route(url string, handler types.PingrateHttpHandler, middlewares ...types.Middleware) {
	s.router.Handle(url, s.newHandler(handler, middlewares...))
}

func (s *PingrateApiServer) newHandler(handler types.PingrateHttpHandler, middlewares ...types.Middleware) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := &types.Ctx{
			Res:  w,
			Req:  r,
			Data: s.dependencies,
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

func (s *PingrateApiServer) Get(name string) interface{} {
	return s.dependencies[name]
}

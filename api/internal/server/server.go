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
	address string
	router  *http.ServeMux
}

func New(address string) (*PingrateApiServer, error) {
	router := http.NewServeMux()

	return &PingrateApiServer{
		address: address,
		router:  router,
	}, nil
}

func (s *PingrateApiServer) Start() {
	v1 := http.NewServeMux()
	v1.Handle("/api/v1/", http.StripPrefix("/api/v1", s.router))

	MuxMiddleWareChain := middleware.NewMiddlewareChain(
		middleware.NewRequestTracingMiddleware,
	)

	server := http.Server{
		Addr:    s.address,
		Handler: MuxMiddleWareChain.ApplyOn(v1),
	}

	slog.Info("Starting the PingrateApiServer at " + s.address)

	err := server.ListenAndServe()

	if err != nil {
		pingrateErr := eris.Wrap(err, "Error starting HTTP server")

		log.Fatal(eris.ToString(pingrateErr, true))
	}
}

func (s *PingrateApiServer) Route(url string, handler http.Handler) {
	s.router.Handle(url, handler)
}

func (s *PingrateApiServer) NewHandler(handler types.PingrateHttpHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			response := types.ServerResponse{
				Ok:    false,
				Error: constants.ServerGenericError,
			}

			if encodeErr := helper.Encode[types.ServerResponse](w, http.StatusInternalServerError, response); encodeErr != nil {
				log.Fatal(eris.ToString(encodeErr, true))
			}

			slog.Error(fmt.Sprintf("%s TraceId=%s", eris.ToString(err, true), r.Header.Get("traceId")))
		}
	}
}

package server

import (
	"encoding/json"
	"errors"
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
			response, serverResponseErr := s.newErrorServerResponse(err)

			if serverResponseErr != nil {
				res := types.ServerResponse{
					Ok:    false,
					Error: constants.ServerGenericError,
				}

				encodeErr := helper.Encode[types.ServerResponse](w, http.StatusInternalServerError, res)

				if encodeErr != nil {
					log.Fatal(eris.ToString(encodeErr, true))
				}

				slog.Error(fmt.Sprintf(
					"%s TraceId=%s",
					eris.ToString(err, true),
					r.Header.Get("traceId"),
				))

				return
			}

			if response.Error != constants.ServerGenericError {
				encodeErr := helper.Encode[types.ServerResponse](w, http.StatusOK, *response)

				if encodeErr != nil {
					log.Fatal(eris.ToString(encodeErr, true))
				}

				return
			}

			encodeErr := helper.Encode[types.ServerResponse](w, http.StatusInternalServerError, *response)

			if encodeErr != nil {
				log.Fatal(eris.ToString(encodeErr, true))
			}

			slog.Error(fmt.Sprintf("%s TraceId=%s", eris.ToString(err, true), r.Header.Get("traceId")))
		}
	}
}

func (s *PingrateApiServer) newErrorServerResponse(err error) (*types.ServerResponse, error) {
	if errors.Is(err, types.ValidationError) {
		strJson, jsonErr := helper.ExtractJSON(err.Error())

		if jsonErr != nil {
			return nil, eris.New(jsonErr.Error())
		}

		var result map[string]string

		unMarshalErr := json.Unmarshal([]byte(strJson), &result)

		if unMarshalErr != nil {
			return nil, eris.New(unMarshalErr.Error())
		}

		return &types.ServerResponse{
			Ok:    false,
			Error: constants.ValidationError,
			Data:  result,
		}, nil
	}

	return &types.ServerResponse{
		Ok:    false,
		Error: constants.ServerGenericError,
	}, nil
}

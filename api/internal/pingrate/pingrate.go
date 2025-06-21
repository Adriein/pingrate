package pingrate

import (
	"github.com/adriein/pingrate/internal/shared/container"
	"net/http"
)

type Pingrate struct {
	api       *HttpServer
	container container.Container
}

func New(address string) *Pingrate {
	router := http.NewServeMux()
	cont := container.New()

	httpServer := &HttpServer{
		address:   address,
		router:    router,
		container: cont,
	}

	return &Pingrate{
		api:       httpServer,
		container: cont,
	}
}

func (p *Pingrate) Get(instance string) interface{} {
	return p.container.Get(instance)
}

func (p *Pingrate) GetHttpServer() *HttpServer {
	return p.api
}

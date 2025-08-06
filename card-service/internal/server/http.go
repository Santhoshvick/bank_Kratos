package server

import (
	v1 "card-service/api/helloworld/v1"
	"card-service/internal/conf"
	"card-service/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	h "github.com/gorilla/handlers"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.CardService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	cors := http.Filter(h.CORS(
        h.AllowedOrigins([]string{"*"}),
        h.AllowedMethods([]string{"GET", "POST", "PATCH", "DELETE","PUT"}),
        h.AllowedHeaders([]string{"Content-Type", "Content-Disposition"}),
        h.ExposedHeaders([]string{"Content-Disposition", "Content-Type"}),
    ))
    opts = append(opts, cors)
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterGreeterHTTPServer(srv, greeter)
	return srv
}

package service

import (
	"github.com/Swapica/rpc-proxy-svc/internal/service/handlers"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxNetworks(s.cfg.Networks()),
		),
	)
	r.Route("/integrations/rpc-proxy", func(r chi.Router) {
		r.Post("/{network}", handlers.RPC)
	})

	return r
}

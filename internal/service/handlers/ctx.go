package handlers

import (
	"context"
	"net/http"

	"github.com/Swapica/rpc-proxy-svc/internal/config"
	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey      ctxKey = iota
	networksCtxKey ctxKey = iota
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxNetworks(n config.Networks) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, networksCtxKey, n)
	}
}

func Networks(r *http.Request) config.Networks {
	return r.Context().Value(networksCtxKey).(config.Networks)
}

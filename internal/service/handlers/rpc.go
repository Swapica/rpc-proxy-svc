package handlers

import (
	"io"
	"net/http"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func RPC(w http.ResponseWriter, r *http.Request) {
	net := chi.URLParam(r, "network")
	endpoint := Networks(r).Get(net)
	if endpoint == nil {
		Log(r).Debugf("network %q not found", net)
		ape.RenderErr(w, problems.NotFound())
		return
	}

	resp, err := http.Post(*endpoint, "application/json", r.Body)
	if err != nil {
		Log(r).WithError(err).Error("failed to send POST request")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		Log(r).WithError(err).Error("failed to read response body")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(resp.StatusCode) // must be here, because it's ignored after w.Write
	if _, err = w.Write(body); err != nil {
		Log(r).WithError(err).Error("failed to write response")
		ape.RenderErr(w, problems.InternalError())
		return
	}
}

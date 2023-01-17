package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func RPC(w http.ResponseWriter, r *http.Request) {
	// TODO: implement RPC logic when config is present
	endpoint := "http://"
	resp, err := http.Post(endpoint, "application/json", r.Body)
	if err != nil {
		Log(r).WithError(err).Error("failed to send POST request")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, resp.Body)
	w.WriteHeader(resp.StatusCode)
}

package api

import (
	"io/ioutil"
	"net/http"

	"github.com/tlentz/d2modmaker/d2mod"
	"github.com/tlentz/d2modmaker/d2mod/config"
)

// Handler returns http.Handler for API endpoint
func RunHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			cfg := config.Parse(body)
			d2mod.Make("~/d2-mod-maker-dist/", cfg)
			w.WriteHeader(500)
			return
		}

		w.WriteHeader(200)
		w.Write(body)
	}
}

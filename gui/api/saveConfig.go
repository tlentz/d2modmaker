package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/tlentz/d2modmaker/internal/d2mod/config"
	"github.com/tlentz/d2modmaker/internal/util"
)

// Handler returns http.Handler for API endpoint
func SaveConfigHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		body, err := ioutil.ReadAll(r.Body)
		util.Check(err)

		cfg := config.Parse(body)
		blob, _ := json.MarshalIndent(cfg, "", "  ")
		ioutil.WriteFile("./cfg.json", blob, 0644) // write json

		w.WriteHeader(200)
		fmt.Println("Writing...")
		util.PP(cfg)
	}
}

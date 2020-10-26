package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tlentz/d2modmaker/internal/d2mod/config"
	"github.com/tlentz/d2modmaker/internal/util"
)

// Handler returns http.Handler for API endpoint
func GetConfigHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		cfg := config.Read("./cfg.json")
		blob, _ := json.Marshal(cfg)
		w.WriteHeader(200)
		w.Write(blob)
		fmt.Println("Fetching...")
		util.PP(cfg)
	}
}

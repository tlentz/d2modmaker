package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tlentz/d2modmaker/internal/d2mod"
	"github.com/tlentz/d2modmaker/internal/d2mod/config"
)

// Handler returns http.Handler for API endpoint
func RunHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(500)
			log.Fatal(err)
			return
		} else {
			cfg := config.Parse(body)
			d2mod.Make("./", cfg)

			w.WriteHeader(200)
			fmt.Println("Done!")
		}
	}
}

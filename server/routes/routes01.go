package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func Routes01(r *chi.Mux) {
	// Create a few simple routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, From K8 Mini App ouo <3")
	})

	r.Get("/pod", func(w http.ResponseWriter, r *http.Request) {
		// Read the pod's name from the Downward API
		fmt.Fprint(w, os.Getenv("POD_NAME"))
	})

	r.Get("/secrets", func(w http.ResponseWriter, r *http.Request) {
		// Get all environment variables
		env := os.Environ()

		// Iterate over the environment variables and print them
		fmt.Println("_________________Env Variables_________________")
		for _, v := range env {
			fmt.Println(v)
		}
		fmt.Println("_________________End_________________")

		// Read the pod's name from the Downward API
		v := map[string]string{
			"secret":       os.Getenv("SECRET"),
			"nestedSecret": os.Getenv("NESTED.SECRET"),
		}

		json.NewEncoder(w).Encode(&v)
	})
}

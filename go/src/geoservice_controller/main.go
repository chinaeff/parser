package main

import (
	"geoservice/controller"
	"geoservice/controller/responder"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"
)

// sfsdfsdf
func main() {
	r := chi.NewRouter()

	responder := &responder.DefaultResponder{} // Замените DefaultResponder на ваш тип Responder.

	geoController := controller.NewGeoController(responder)

	r.Route("/api/address", func(r chi.Router) {
		r.Get("/search", geoController.SearchHandler)
		r.Get("/geocode", geoController.GeocodeHandler)
	})

	http.ListenAndServe(":8080", r)
}

type ReverseProxy struct {
	host string
	port string
}

func NewReverseProxy(host, port string) *ReverseProxy {
	return &ReverseProxy{
		host: host,
		port: port,
	}
}

func (rp *ReverseProxy) ReverseProxy(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/" || r.URL.Path == "/api" {
			w.Write([]byte("Hello from API"))
		} else {
			targetURL := "http://" + rp.host + ":" + rp.port
			target, _ := url.Parse(targetURL)

			proxy := httputil.NewSingleHostReverseProxy(target)
			proxy.ServeHTTP(w, r)
		}
	})
}

const content = ``

func WorkerTest() {
	t := time.NewTicker(1 * time.Second)
	var b byte = 0
	for {
		select {
		case <-t.C:
			err := os.WriteFile("/app/content/_index.md", []byte(content), 0644)
			if err != nil {
				log.Println(err)
			}
			b++
		}
	}
}

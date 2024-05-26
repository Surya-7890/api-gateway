package api

import (
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Server struct {
	Address string
}

func NewServer(address string) *Server {
	return &Server{
		Address: address,
	}
}

func (s *Server) StartServer() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello")
	})

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		proxy := httputil.NewSingleHostReverseProxy(&url.URL{
			Host:   "localhost:7000",
			Scheme: "http",
		})
		r.Header.Set("X-Email", "")
		r.Header.Set("X-Role", "")
		proxy.ServeHTTP(w, r)
	})
	http.ListenAndServe(s.Address, nil)
}

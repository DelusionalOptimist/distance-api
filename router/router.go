package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DelusionalOptimist/distance-api/handlers"
)

type Router struct {
	mux *http.ServeMux
	port int
}

func NewRouter(port int, logger *log.Logger) *Router {
	mux := http.NewServeMux()

	handler := handlers.NewHandler(logger)
	mux.HandleFunc("/getdistance", handler.GetDistance)

	return &Router{
		mux:  mux,
		port: port,
	}
}

func (r *Router) Run() error {
	err := http.ListenAndServe(fmt.Sprintf(":%d", r.port), r.mux)
	if err != nil {
		return err
	}
	return nil
}

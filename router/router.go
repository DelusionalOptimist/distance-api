package router

import (
	"log"
	"net/http"

	"github.com/DelusionalOptimist/distance-api/handlers"
	"github.com/DelusionalOptimist/distance-api/utils"
)

type Router struct {
	mux *http.ServeMux
}

func NewRouter(logger *log.Logger) *Router {
	mux := http.NewServeMux()

	handler := handlers.NewHandler(logger)

	mux.Handle("/", http.FileServer(http.Dir("./static")))
	mux.HandleFunc("/getDistance", handler.GetDistance)

	return &Router{
		mux:  mux,
	}
}

func (r *Router) Run() error {
	err := http.ListenAndServe(":" + utils.GlobalConfig.Port, r.mux)
	if err != nil {
		return err
	}
	return nil
}

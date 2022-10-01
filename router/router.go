package router

import (
	"log"
	"net/http"

	"github.com/DelusionalOptimist/distance-api/handlers"
	"github.com/DelusionalOptimist/distance-api/utils"
)

// Router object
type Router struct {
	mux *http.ServeMux
}

// creates paths
func NewRouter(logger *log.Logger) *Router {
	mux := http.NewServeMux()

	handler := handlers.NewHandler(logger)

	// serve static content
	mux.Handle("/", http.FileServer(http.Dir("./static")))

	// endpoint for the distance API
	mux.HandleFunc("/getDistance", handler.GetDistance)

	return &Router{
		mux: mux,
	}
}

// run the router with the set configuration
func (r *Router) Run() error {
	err := http.ListenAndServe(":"+utils.GlobalConfig.Port, r.mux)
	if err != nil {
		return err
	}
	return nil
}

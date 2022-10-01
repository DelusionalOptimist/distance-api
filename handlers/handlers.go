package handlers

import "log"

// Handler object
type Handler struct {
	Log *log.Logger
}

// Creates a new handler
func NewHandler(logger *log.Logger) *Handler {
	return &Handler{
		Log: logger,
	}
}

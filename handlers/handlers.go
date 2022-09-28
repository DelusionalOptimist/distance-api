package handlers

import "log"

type Handler struct {
	Log *log.Logger
}

func NewHandler(logger *log.Logger) *Handler {
	return &Handler{
		Log: logger,
	}
}

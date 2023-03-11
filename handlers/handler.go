package handlers

import (
	"github.com/aabri-assignments/assignment-golang/pkg/decoder"
	"github.com/aabri-assignments/assignment-golang/pkg/logger"
)

type Handler struct {
	dec decoder.Decoder
}

func NewHandler() *Handler {
	log := logger.InitLogger()

	return &Handler{
		dec: decoder.New(log),
	}
}

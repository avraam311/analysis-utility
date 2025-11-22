package analysis

import (
	"context"

	"github.com/avraam311/analysis-utility/internal/models/domain"
)

type Service interface {
	GetAnalysis(ctx context.Context) (*domain.Analysis, error)
}

type Handler struct {
	service Service
}

func New(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

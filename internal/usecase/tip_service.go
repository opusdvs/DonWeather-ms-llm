package usecase

import (
	"context"

	"github.com/opusdvs/DonWeather-ms-ollama/internal/domain"
)

type TipService struct {
	tipProvider domain.TipProvider
}

func NewTipService(tipProvider domain.TipProvider) *TipService {
	return &TipService{tipProvider: tipProvider}
}

func (s *TipService) GetTip(ctx context.Context, prediction domain.Prediction) (*domain.Tip, error) {
	return s.tipProvider.GetTip(ctx, prediction)
}

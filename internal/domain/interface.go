package domain

import (
	"context"
)

type TipProvider interface {
	GetTip(ctx context.Context, prediction Prediction) (*Tip, error)
}

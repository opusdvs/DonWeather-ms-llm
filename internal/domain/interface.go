package domain

import (
	"context"
)

type TipProvider interface {
	GetTip(ctx context.Context, input Prediction) (*Tip, error)
}

package domain

import (
	"clean-archy/internal/entity"
	"context"
)

type ReportUsecase interface {
	TotalGameSales(ctx context.Context) ([]entity.TotalGameSales, error)

	MostPopularGame(ctx context.Context) ([]entity.MostPopularGame, error)

	RevenuePerGame(ctx context.Context) ([]entity.RevenuePerGame, error)

	PlayerCountPerGame(ctx context.Context) ([]entity.PlayerCountPerGame, error)
}

package domain

import (
	"clean-archy/internal/entity"
	"context"
)

type ReportRepository interface {
	GetTotalGameSales(ctx context.Context) ([]entity.TotalGameSales, error)

	GetMostPopularGame(ctx context.Context) ([]entity.MostPopularGame, error)

	GetRevenuePerGame(ctx context.Context) ([]entity.RevenuePerGame, error)

	GetPlayerCountPerGame(ctx context.Context) ([]entity.PlayerCountPerGame, error)
}
